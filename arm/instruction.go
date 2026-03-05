package arm

import (
	"fmt"
	"golite/ir"
	"strings"
)

func isGlobal(name string) bool {
	if name == "" {
		return false
	}
	if name[0] == 'x' || name[0] == 'w' || name[0] == '#' {
		return false
	}
	return true
}

func (b *builder) loadGlobalAddress(dst, global string) {
	if b.isApple {
		b.emitText(fmt.Sprintf("\tadrp\t%s, %s@PAGE", dst, global))
		b.emitText(fmt.Sprintf("\tadd\t%s, %s, %s@PAGEOFF", dst, dst, global))
	} else {
		b.emitText(fmt.Sprintf("\tadrp\t%s, %s", dst, global))
		b.emitText(fmt.Sprintf("\tadd\t%s, %s, :lo12:%s", dst, dst, global))
	}
}

func (b *builder) operandToARM(val ir.Value, fn *ir.Function) string {
	return b.operandToARMScratch(val, fn, "x9")
}

func (b *builder) operandToARMScratch(val ir.Value, fn *ir.Function, scratch string) string {
	switch v := val.(type) {
	case *ir.Register:

		phys := v.AllocatedPhyReg()

		if strings.HasPrefix(phys, "[") && strings.HasSuffix(phys, "]") {
			phys = phys[1 : len(phys)-1]
		}

		if phys == "" {
			return scratch
		}
		return phys

	default:
		name := val.String()
		if strings.HasPrefix(name, "%") {
			rawName := strings.TrimPrefix(name, "%")
			for i, param := range fn.Entry.Params {
				if param.Name == rawName {
					offset := (i + 1) * 8
					b.emitText(fmt.Sprintf("\tldr\t%s, [x29, #-%d]", scratch, offset))
					return scratch
				}
			}

			var id int
			_, err := fmt.Sscanf(rawName, "%d", &id)
			if err == nil {
				offset := (id + 1) * 8
				b.emitText(fmt.Sprintf("\tldr\t%s, [x29, #-%d]", scratch, offset))
				return scratch
			}
		}

		if strings.HasPrefix(name, "@") {
			raw := strings.TrimPrefix(name, "@")

			return raw
		}

		if name == "null" || name == "false" || name == "zeroinitializer" {
			return "#0"
		}
		if name == "true" {
			return "#1"
		}

		if name[0] >= '0' && name[0] <= '9' || name[0] == '-' {
			return "#" + name
		}

		return name
	}
}

func (b *builder) storeToARM(val ir.Value, fn *ir.Function, src string) {
	switch v := val.(type) {
	case *ir.Register:
		phys := v.AllocatedPhyReg()
		if strings.HasPrefix(phys, "[") && strings.HasSuffix(phys, "]") {
			phys = phys[1 : len(phys)-1]
		}

		if phys == "" {
			name := val.String()
			if strings.HasPrefix(name, "%") {
				rawName := strings.TrimPrefix(name, "%")
				for i, param := range fn.Entry.Params {
					if param.Name == rawName {
						offset := (i + 1) * 8
						b.emitText(fmt.Sprintf("\tstr\t%s, [x29, #-%d]", src, offset))
						return
					}
				}

				var id int
				_, err := fmt.Sscanf(rawName, "%d", &id)
				if err == nil {
					offset := (id + 1) * 8
					b.emitText(fmt.Sprintf("\tstr\t%s, [x29, #-%d]", src, offset))
					return
				}
			}
			return
		}
		b.emitText(fmt.Sprintf("\tmov\t%s, %s", phys, src))
	}
}

func (b *builder) translateInstruction(inst ir.Instruction, fn *ir.Function) {
	switch v := inst.(type) {

	case *ir.BinaryOp:
		dst := "x9"
		if v.Result != nil {
			phys := v.Result.AllocatedPhyReg()
			if strings.HasPrefix(phys, "[") && strings.HasSuffix(phys, "]") {
				phys = phys[1 : len(phys)-1]
			}
			if phys != "" {
				dst = phys
			}
		}

		left := b.operandToARMScratch(v.Left, fn, "x9")
		if left == dst && dst == "x9" {
			left = b.operandToARMScratch(v.Left, fn, "x11")
		}

		right := b.operandToARMScratch(v.Right, fn, "x10")

		op := "add"
		switch v.Op {
		case "add":
			op = "add"
		case "sub":
			op = "sub"
		case "mul":
			op = "mul"
		case "sdiv":
			op = "sdiv"
		case "srem":
			op = "srem"
		case "and":
			op = "and"
		case "or":
			op = "orr"
		case "xor":
			op = "eor"
		}

		if left[0] == '#' {
			b.emitText(fmt.Sprintf("\tmov\t%s, %s", "x11", left))
			left = "x11"
		}

		if right[0] == '#' && (op == "mul" || op == "sdiv" || len(right) > 4) {
			b.emitText(fmt.Sprintf("\tmov\tx10, %s", right))
			right = "x10"
		}

		b.emitText(fmt.Sprintf("\t%s\t%s, %s, %s", op, dst, left, right))
		b.storeToARM(v.Result, fn, dst)

	case *ir.Icmp:
		left := b.operandToARMScratch(v.Left, fn, "x9")
		right := b.operandToARMScratch(v.Right, fn, "x10")

		dst := "x9"
		if v.Result != nil {
			phys := v.Result.AllocatedPhyReg()
			if strings.HasPrefix(phys, "[") && strings.HasSuffix(phys, "]") {
				phys = phys[1 : len(phys)-1]
			}
			if phys != "" {
				dst = phys
			}
		}

		if left[0] == '#' {
			b.emitText(fmt.Sprintf("\tmov\tx11, %s", left))
			left = "x11"
		}
		b.emitText(fmt.Sprintf("\tcmp\t%s, %s", left, right))
		cond := "eq"
		switch v.Cond {
		case "eq":
			cond = "eq"
		case "ne":
			cond = "ne"
		case "sgt":
			cond = "gt"
		case "sge":
			cond = "ge"
		case "slt":
			cond = "lt"
		case "sle":
			cond = "le"
		}

		b.emitText(fmt.Sprintf("\tcset\t%s, %s", dst, cond))
		b.storeToARM(v.Result, fn, dst)

	case *ir.Branch:
		target := fmt.Sprintf(".L%s_%s", fn.Entry.Name, v.Target)
		b.emitText(fmt.Sprintf("\tb\t%s", target))

	case *ir.CondBranch:
		cond := b.operandToARM(v.Cond, fn)

		if cond[0] == '#' {
			b.emitText(fmt.Sprintf("\tmov\tx10, %s", cond))
			cond = "x10"
		}

		b.emitText(fmt.Sprintf("\tcmp\t%s, #0", cond))
		tTarget := fmt.Sprintf(".L%s_%s", fn.Entry.Name, v.TrueL)
		fTarget := fmt.Sprintf(".L%s_%s", fn.Entry.Name, v.FalseL)
		b.emitText(fmt.Sprintf("\tb.ne\t%s", tTarget))
		b.emitText(fmt.Sprintf("\tb\t%s", fTarget))

	case *ir.Call:
		variadicSpace := 0
		if b.isApple && v.IsVariadic && len(v.Args) > 1 {
			variadicCount := len(v.Args) - 1
			variadicSpace = (variadicCount*8 + 15) & ^15
			b.emitText(fmt.Sprintf("\tsub\tsp, sp, #%d", variadicSpace))
		}

		for i, arg := range v.Args {
			argARM := b.operandToARM(arg, fn)

			if b.isApple && v.IsVariadic && i > 0 {
				targetTemp := "x9"
				if isGlobal(argARM) {
					b.loadGlobalAddress(targetTemp, argARM)
				} else if argARM[0] == '#' {
					b.emitText(fmt.Sprintf("\tmov\t%s, %s", targetTemp, argARM))
				} else {
					targetTemp = argARM
				}
				offset := (i - 1) * 8
				b.emitText(fmt.Sprintf("\tstr\t%s, [sp, #%d]", targetTemp, offset))
			} else {
				target := fmt.Sprintf("x%d", i)
				if isGlobal(argARM) {
					b.loadGlobalAddress(target, argARM)
				} else if argARM != target {
					b.emitText(fmt.Sprintf("\tmov\t%s, %s", target, argARM))
				}
			}
		}

		name := v.Name
		if b.isApple {
			name = "_" + name
		}
		b.emitText(fmt.Sprintf("\tbl\t%s", name))

		if variadicSpace > 0 {
			b.emitText(fmt.Sprintf("\tadd\tsp, sp, #%d", variadicSpace))
		}

		if v.Result != nil {
			b.storeToARM(v.Result, fn, "x0")
		}

	case *ir.Store:
		src := b.operandToARM(v.Src, fn)
		dst := b.operandToARM(v.Dst, fn)

		if src[0] == '#' {
			b.emitText(fmt.Sprintf("\tmov\tx10, %s", src))
			src = "x10"
		} else if isGlobal(src) {
			b.loadGlobalAddress("x10", src)
			src = "x10"
		}

		if isGlobal(dst) {
			b.loadGlobalAddress("x9", dst)
			b.emitText(fmt.Sprintf("\tstr\t%s, [x9]", src))
		} else {
			b.emitText(fmt.Sprintf("\tstr\t%s, [%s]", src, dst))
		}

	case *ir.Load:
		src := b.operandToARM(v.Src, fn)
		dst := "x9"
		if v.Result != nil {
			phys := v.Result.AllocatedPhyReg()
			if strings.HasPrefix(phys, "[") && strings.HasSuffix(phys, "]") {
				phys = phys[1 : len(phys)-1]
			}
			if phys != "" {
				dst = phys
			}
		}

		if isGlobal(src) {
			b.loadGlobalAddress("x10", src)
			b.emitText(fmt.Sprintf("\tldr\t%s, [x10]", dst))
		} else {
			b.emitText(fmt.Sprintf("\tldr\t%s, [%s]", dst, src))
		}
		b.storeToARM(v.Result, fn, dst)

	case *ir.Mov:
		src := b.operandToARM(v.Src, fn)
		dst := "x9"
		if v.Result != nil {
			phys := v.Result.AllocatedPhyReg()
			if strings.HasPrefix(phys, "[") && strings.HasSuffix(phys, "]") {
				phys = phys[1 : len(phys)-1]
			}
			if phys != "" {
				dst = phys
			}
		}

		if isGlobal(src) {
			b.loadGlobalAddress(dst, src)
		} else if dst != src {
			b.emitText(fmt.Sprintf("\tmov\t%s, %s", dst, src))
		}
		b.storeToARM(v.Result, fn, dst)

	case *ir.Bitcast:
		src := b.operandToARM(v.Src, fn)
		dst := "x9"
		if v.Result != nil {
			phys := v.Result.AllocatedPhyReg()
			if strings.HasPrefix(phys, "[") && strings.HasSuffix(phys, "]") {
				phys = phys[1 : len(phys)-1]
			}
			if phys != "" {
				dst = phys
			}
		}

		if isGlobal(src) {
			b.loadGlobalAddress(dst, src)
		} else if dst != src {
			b.emitText(fmt.Sprintf("\tmov\t%s, %s", dst, src))
		}
		b.storeToARM(v.Result, fn, dst)

	case *ir.Gep:
		ptr := b.operandToARM(v.Ptr, fn)
		dst := "x9"
		if v.Result != nil {
			phys := v.Result.AllocatedPhyReg()
			if strings.HasPrefix(phys, "[") && strings.HasSuffix(phys, "]") {
				phys = phys[1 : len(phys)-1]
			}
			if phys != "" {
				dst = phys
			}
		}
		offset := v.Index * 8

		if isGlobal(ptr) {
			b.loadGlobalAddress("x10", ptr)
			ptr = "x10"
		}

		if offset == 0 {
			if dst != ptr {
				b.emitText(fmt.Sprintf("\tmov\t%s, %s", dst, ptr))
			}
		} else {
			b.emitText(fmt.Sprintf("\tadd\t%s, %s, #%d", dst, ptr, offset))
		}
		b.storeToARM(v.Result, fn, dst)

	case *ir.Alloca:
		if v.Result != nil {
			dst := b.operandToARM(v.Result, fn)
			offset := (v.Result.ID + 1) * 8
			b.emitText(fmt.Sprintf("\tsub\t%s, x29, #%d", dst, offset))
		}
	case *ir.Return:
		if v.Val != nil {
			retPhys := b.operandToARM(v.Val, fn)
			outSize := "x0"
			if isGlobal(retPhys) {
				b.loadGlobalAddress("x0", retPhys)
			} else if retPhys[0] == '#' {
				b.emitText(fmt.Sprintf("\tmov\t%s, %s", outSize, retPhys))
			} else if retPhys != outSize {
				b.emitText(fmt.Sprintf("\tmov\t%s, %s", outSize, retPhys))
			}
		}
		b.emitText(fmt.Sprintf("\tb\t.L%s_end", fn.Entry.Name))

	default:
		b.emitText(fmt.Sprintf("\t// UNIMPLEMENTED: %s", v.String()))
	}
}
