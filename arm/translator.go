package arm

import (
	"fmt"
	"golite/ir"
	"strings"
)

func (b *builder) translateGlobals() {
	for _, g := range b.irBuilder.Globals() {
		name := g.Name

		if !strings.HasPrefix(name, ".") {

		}

		val := g.Value
		if val == "" || val == "0" || val == "null" || val == "false" || val == "zeroinitializer" {
			if !g.IsPrivate {
				b.emitBss(fmt.Sprintf("\t.globl\t%s", name))
			}
			b.emitBss("\t.p2align\t3")
			b.emitBss(fmt.Sprintf("%s:", name))
			b.emitBss("\t.space\t8")
		} else if strings.HasPrefix(val, "c\"") {
			s := strings.TrimPrefix(val, "c\"")
			s = strings.TrimSuffix(s, "\"")

			s = strings.ReplaceAll(s, "\\0A", "\\n")
			s = strings.ReplaceAll(s, "\\00", "")

			b.emitData(fmt.Sprintf("%s:", name))
			b.emitData(fmt.Sprintf("\t.asciz\t\"%s\"", s))
		}
	}
}

func (b *builder) translateFunctions() {
	for _, fn := range b.irBuilder.Functions() {
		name := fn.Entry.Name
		if b.isApple {
			name = "_" + name
		}

		b.emitText(fmt.Sprintf("\t.globl\t%s", name))
		b.emitText("\t.p2align\t2")
		b.emitText(fmt.Sprintf("%s:", name))

		b.emitText("\tstp\tx29, x30, [sp, #-16]!")
		b.emitText("\tmov\tx29, sp")

		paramCount := len(fn.Entry.Params)
		maxReg := paramCount
		for _, block := range fn.Blocks {
			for _, inst := range block.Instructions {
				switch i := inst.(type) {
				case *ir.Load:
					if i.Result != nil && i.Result.ID > maxReg {
						maxReg = i.Result.ID
					}
				case *ir.Store:
					if i.Dst != nil {
						if r, ok := i.Dst.(*ir.Register); ok && r.ID > maxReg {
							maxReg = r.ID
						}
					}
				case *ir.BinaryOp:
					if i.Result != nil && i.Result.ID > maxReg {
						maxReg = i.Result.ID
					}
				case *ir.Icmp:
					if i.Result != nil && i.Result.ID > maxReg {
						maxReg = i.Result.ID
					}
				case *ir.Call:
					if i.Result != nil {
						if i.Result.ID > maxReg {
							maxReg = i.Result.ID
						}
					}
				case *ir.Gep:
					if i.Result != nil && i.Result.ID > maxReg {
						maxReg = i.Result.ID
					}
				case *ir.Bitcast:
					if i.Result != nil && i.Result.ID > maxReg {
						maxReg = i.Result.ID
					}
				case *ir.Mov:
					if i.Result != nil && i.Result.ID > maxReg {
						maxReg = i.Result.ID
					}
				case *ir.Alloca:
					if i.Result != nil && i.Result.ID > maxReg {
						maxReg = i.Result.ID
					}
				}
			}
		}

		if maxReg >= 0 {
			stackSpace := ((maxReg * 8) + 15) & ^15
			if stackSpace > 0 {
				b.emitText(fmt.Sprintf("\tsub\tsp, sp, #%d", stackSpace))
			}
			for i := 0; i < paramCount && i < 8; i++ {
				offset := (i + 1) * 8
				b.emitText(fmt.Sprintf("\tstr\tx%d, [x29, #-%d]", i, offset))
			}
		}

		b.emitText("\tstp\tx19, x20, [sp, #-16]!")
		b.emitText("\tstp\tx21, x22, [sp, #-16]!")
		b.emitText("\tstp\tx23, x24, [sp, #-16]!")
		b.emitText("\tstp\tx25, x26, [sp, #-16]!")
		b.emitText("\tstp\tx27, x28, [sp, #-16]!")

		for _, block := range fn.Blocks {

			b.emitText(fmt.Sprintf(".L%s_%s:", fn.Entry.Name, block.Label))
			for _, inst := range block.Instructions {
				b.translateInstruction(inst, fn)
			}
		}

		b.emitText(fmt.Sprintf(".L%s_end:", fn.Entry.Name))
		b.emitText("\tldp\tx27, x28, [sp], #16")
		b.emitText("\tldp\tx25, x26, [sp], #16")
		b.emitText("\tldp\tx23, x24, [sp], #16")
		b.emitText("\tldp\tx21, x22, [sp], #16")
		b.emitText("\tldp\tx19, x20, [sp], #16")
		b.emitText("\tmov\tsp, x29")
		b.emitText("\tldp\tx29, x30, [sp], #16")
		b.emitText("\tret")
	}
}
