package ir

import (
	"fmt"
	"golite/types"
)

func (b *builder) Mem2RegPass() {

	b.addReadScratchGlobal()

	for _, fn := range b.functions {
		mem2regFunction(fn, b)
	}
}

func (b *builder) addReadScratchGlobal() {

	for _, g := range b.globals {
		if g.Name == ".read_scratch" {
			return
		}
	}
	b.globals = append(b.globals, &Global{
		Name:    ".read_scratch",
		Ty:      types.IntTySig,
		Value:   "0",
		IsConst: false,
	})
}

type scanfTarget struct {
	allocaRegID int
	loadReg     *Register
}

func mem2regFunction(fn *Function, b *builder) {
	if len(fn.Blocks) == 0 {
		return
	}

	advanceRegisterID(fn, b)

	buildCFG(fn)

	promotable := identifyPromotable(fn)
	if len(promotable) == 0 {
		return
	}

	scanfTargets := prescanScanfTargets(fn, promotable, b)

	inMap, phiMap := ssaDataflow(fn, promotable, scanfTargets, b)

	rewriteFunction(fn, promotable, inMap, phiMap, scanfTargets, b)
}

func buildCFG(fn *Function) {
	blockMap := make(map[string]*BasicBlock)
	for _, block := range fn.Blocks {
		block.Predecessors = nil
		block.Successors = nil
		blockMap[block.Label] = block
	}

	for _, block := range fn.Blocks {
		if block.Terminator == nil {
			continue
		}
		switch t := block.Terminator.(type) {
		case *Branch:
			if succ, ok := blockMap[t.Target]; ok {
				block.Successors = append(block.Successors, succ)
				succ.Predecessors = append(succ.Predecessors, block)
			}
		case *CondBranch:
			if succ, ok := blockMap[t.TrueL]; ok {
				block.Successors = append(block.Successors, succ)
				succ.Predecessors = append(succ.Predecessors, block)
			}
			if succ, ok := blockMap[t.FalseL]; ok {
				block.Successors = append(block.Successors, succ)
				succ.Predecessors = append(succ.Predecessors, block)
			}
		}
	}
}

func identifyPromotable(fn *Function) map[int]string {
	promotable := make(map[int]string)
	if len(fn.Blocks) == 0 {
		return promotable
	}

	entryBlock := fn.Blocks[0]

	allocaRegs := make(map[int]types.Type)
	for _, inst := range entryBlock.Instructions {
		if alloca, ok := inst.(*Alloca); ok {
			allocaRegs[alloca.Result.ID] = alloca.Ty
		}
	}

	gepBases := make(map[int]bool)
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if gep, ok := inst.(*Gep); ok {
				if reg, ok := gep.Ptr.(*Register); ok {
					gepBases[reg.ID] = true
				}
			}
		}
	}

	for id := range allocaRegs {
		if gepBases[id] {
			continue
		}
		promotable[id] = fmt.Sprintf("t%d", id)
	}

	return promotable
}

func prescanScanfTargets(fn *Function, promotable map[int]string, b *builder) map[string][]scanfTarget {
	targets := make(map[string][]scanfTarget)

	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if call, ok := inst.(*Call); ok {
				if call.Name == "scanf" && len(call.Args) >= 2 {
					lastArg := call.Args[len(call.Args)-1]
					if reg, ok := lastArg.(*Register); ok {
						if _, isPromotable := promotable[reg.ID]; isPromotable {
							ty := findAllocaTypeInFunction(fn, reg.ID)
							loadReg := &Register{ID: b.registerID, Ty: ty}
							b.registerID++
							targets[block.Label] = append(targets[block.Label], scanfTarget{
								allocaRegID: reg.ID,
								loadReg:     loadReg,
							})
						}
					}
				}
			}
		}
	}

	return targets
}

func ssaDataflow(fn *Function, promotable map[int]string,
	scanfTargets map[string][]scanfTarget, b *builder) (
	map[string]map[int]Value, map[string]map[int]*Phi) {

	outMap := make(map[string]map[int]Value)
	inMap := make(map[string]map[int]Value)
	phiMap := make(map[string]map[int]*Phi)

	for _, block := range fn.Blocks {
		outMap[block.Label] = make(map[int]Value)
		inMap[block.Label] = make(map[int]Value)
		phiMap[block.Label] = make(map[int]*Phi)
	}

	phiRegs := make(map[string]map[int]*Register)
	for _, block := range fn.Blocks {
		phiRegs[block.Label] = make(map[int]*Register)
	}

	changed := true
	for changed {
		changed = false

		for _, block := range fn.Blocks {
			oldOut := copyValueMap(outMap[block.Label])

			if block == fn.Blocks[0] {

				for regID := range promotable {
					if _, exists := inMap[block.Label][regID]; !exists {
						ty := findAllocaType(block, regID)
						inMap[block.Label][regID] = defaultValue(ty)
					}
				}
			} else {

				for regID := range promotable {
					vals := make([]Value, 0)
					predLabels := make([]string, 0)

					for _, pred := range block.Predecessors {
						if v, ok := outMap[pred.Label][regID]; ok {
							vals = append(vals, v)
							predLabels = append(predLabels, pred.Label)
						}
					}

					if len(vals) == 0 {
						continue
					}

					allSame := true
					for i := 1; i < len(vals); i++ {
						if vals[i].String() != vals[0].String() {
							allSame = false
							break
						}
					}

					if allSame && phiRegs[block.Label][regID] == nil {

						inMap[block.Label][regID] = vals[0]
					} else {

						phiReg := phiRegs[block.Label][regID]
						if phiReg == nil {

							ty := findAllocaTypeInFunction(fn, regID)
							phiReg = &Register{ID: b.registerID, Ty: ty}
							b.registerID++
							phiRegs[block.Label][regID] = phiReg
						}

						newEntries := make([]PhiEntry, len(vals))
						for i := range vals {
							newEntries[i] = PhiEntry{Value: vals[i], Block: predLabels[i]}
						}

						existingPhi, hasPhi := phiMap[block.Label][regID]
						if !hasPhi {
							phi := &Phi{
								Result:  phiReg,
								Entries: newEntries,
							}
							phiMap[block.Label][regID] = phi
							changed = true
						} else {
							if !phiEntriesEqual(existingPhi.Entries, newEntries) {
								existingPhi.Entries = newEntries
								changed = true
							}
						}
						inMap[block.Label][regID] = phiReg
					}
				}
			}

			current := copyValueMap(inMap[block.Label])
			valueOf := make(map[int]Value)

			scanfIdx := 0

			for _, inst := range block.Instructions {
				switch i := inst.(type) {
				case *Load:
					if srcReg, ok := i.Src.(*Register); ok {
						if _, isPromotable := promotable[srcReg.ID]; isPromotable {
							if val, exists := current[srcReg.ID]; exists {
								valueOf[i.Result.ID] = val
							}
							continue
						}
					}
					valueOf[i.Result.ID] = i.Result

				case *Store:
					if dstReg, ok := i.Dst.(*Register); ok {
						if _, isPromotable := promotable[dstReg.ID]; isPromotable {
							storedVal := resolveValue(i.Src, valueOf)
							current[dstReg.ID] = storedVal
							continue
						}
					}

				case *Alloca:

				case *BinaryOp:
					valueOf[i.Result.ID] = i.Result

				case *Icmp:
					valueOf[i.Result.ID] = i.Result

				case *Call:
					if i.Result != nil {
						valueOf[i.Result.ID] = i.Result
					}

					if i.Name == "scanf" && len(i.Args) >= 2 {
						lastArg := i.Args[len(i.Args)-1]
						if reg, ok := lastArg.(*Register); ok {
							if _, isPromotable := promotable[reg.ID]; isPromotable {

								if scanfIdx < len(scanfTargets[block.Label]) {
									st := scanfTargets[block.Label][scanfIdx]
									current[st.allocaRegID] = st.loadReg
									valueOf[st.loadReg.ID] = st.loadReg
									scanfIdx++
								}
							}
						}
					}

				case *Gep:
					valueOf[i.Result.ID] = i.Result

				case *Bitcast:
					valueOf[i.Result.ID] = i.Result
				}
			}

			outMap[block.Label] = current
			if !valueMapsEqual(outMap[block.Label], oldOut) {
				changed = true
			}
		}
	}

	return inMap, phiMap
}

func rewriteFunction(fn *Function, promotable map[int]string,
	inMap map[string]map[int]Value, phiMap map[string]map[int]*Phi,
	scanfTargets map[string][]scanfTarget, b *builder) {

	for _, block := range fn.Blocks {
		current := copyValueMap(inMap[block.Label])
		valueOf := make(map[int]Value)

		newInstructions := make([]Instruction, 0)

		for _, phi := range phiMap[block.Label] {
			newInstructions = append(newInstructions, phi)
		}

		scanfIdx := 0

		for _, inst := range block.Instructions {
			switch i := inst.(type) {
			case *Alloca:
				if _, isPromotable := promotable[i.Result.ID]; isPromotable {
					continue
				}
				newInstructions = append(newInstructions, inst)

			case *Load:
				if srcReg, ok := i.Src.(*Register); ok {
					if _, isPromotable := promotable[srcReg.ID]; isPromotable {
						if val, exists := current[srcReg.ID]; exists {
							valueOf[i.Result.ID] = val
						}
						continue
					}
				}
				i.Src = resolveValue(i.Src, valueOf)
				valueOf[i.Result.ID] = i.Result
				newInstructions = append(newInstructions, i)

			case *Store:
				if dstReg, ok := i.Dst.(*Register); ok {
					if _, isPromotable := promotable[dstReg.ID]; isPromotable {
						storedVal := resolveValue(i.Src, valueOf)
						current[dstReg.ID] = storedVal
						continue
					}
				}
				i.Src = resolveValue(i.Src, valueOf)
				i.Dst = resolveValue(i.Dst, valueOf)
				newInstructions = append(newInstructions, i)

			case *BinaryOp:
				i.Left = resolveValue(i.Left, valueOf)
				i.Right = resolveValue(i.Right, valueOf)
				valueOf[i.Result.ID] = i.Result
				newInstructions = append(newInstructions, i)

			case *Icmp:
				i.Left = resolveValue(i.Left, valueOf)
				i.Right = resolveValue(i.Right, valueOf)
				valueOf[i.Result.ID] = i.Result
				newInstructions = append(newInstructions, i)

			case *Call:

				if i.Name == "scanf" && len(i.Args) >= 2 {
					lastArg := i.Args[len(i.Args)-1]
					if reg, ok := lastArg.(*Register); ok {
						if _, isPromotable := promotable[reg.ID]; isPromotable {

							newArgs := make([]Value, len(i.Args))
							copy(newArgs, i.Args)
							newArgs[0] = resolveValue(newArgs[0], valueOf)
							newArgs[len(newArgs)-1] = &Constant{
								Val: "@.read_scratch",
								Ty:  &PointerType{Base: types.IntTySig},
							}
							i.Args = newArgs
							newInstructions = append(newInstructions, i)

							if scanfIdx < len(scanfTargets[block.Label]) {
								st := scanfTargets[block.Label][scanfIdx]
								loadInst := &Load{
									Result: st.loadReg,
									Src: &Constant{
										Val: "@.read_scratch",
										Ty:  &PointerType{Base: types.IntTySig},
									},
								}
								newInstructions = append(newInstructions, loadInst)
								valueOf[st.loadReg.ID] = st.loadReg
								current[st.allocaRegID] = st.loadReg
								scanfIdx++
							}
							continue
						}
					}
				}

				for j := range i.Args {
					i.Args[j] = resolveValue(i.Args[j], valueOf)
				}
				if i.Result != nil {
					valueOf[i.Result.ID] = i.Result
				}
				newInstructions = append(newInstructions, i)

			case *Gep:
				i.Ptr = resolveValue(i.Ptr, valueOf)
				valueOf[i.Result.ID] = i.Result
				newInstructions = append(newInstructions, i)

			case *Bitcast:
				i.Src = resolveValue(i.Src, valueOf)
				valueOf[i.Result.ID] = i.Result
				newInstructions = append(newInstructions, i)

			case *Return:
				if i.Val != nil {
					i.Val = resolveValue(i.Val, valueOf)
				}
				newInstructions = append(newInstructions, i)

			case *Branch:
				newInstructions = append(newInstructions, i)

			case *CondBranch:
				i.Cond = resolveValue(i.Cond, valueOf)
				newInstructions = append(newInstructions, i)

			default:
				newInstructions = append(newInstructions, inst)
			}
		}

		block.Instructions = newInstructions

		block.Terminator = nil
		for _, inst := range block.Instructions {
			switch inst.(type) {
			case *Return, *Branch, *CondBranch:
				block.Terminator = inst
			}
		}
	}
}

func resolveValue(v Value, valueOf map[int]Value) Value {
	if reg, ok := v.(*Register); ok {
		if resolved, exists := valueOf[reg.ID]; exists {
			if resolvedReg, isReg := resolved.(*Register); isReg && resolvedReg.ID != reg.ID {
				return resolveValue(resolved, valueOf)
			}
			return resolved
		}
	}
	return v
}

func defaultValue(ty types.Type) Value {
	switch LLVMType(ty) {
	case "i64", "i1":
		return &Constant{Val: "0", Ty: ty}
	default:
		return &Constant{Val: "null", Ty: ty}
	}
}

func findAllocaType(block *BasicBlock, regID int) types.Type {
	for _, inst := range block.Instructions {
		if alloca, ok := inst.(*Alloca); ok {
			if alloca.Result.ID == regID {
				return alloca.Ty
			}
		}
	}
	return types.IntTySig
}

func findAllocaTypeInFunction(fn *Function, regID int) types.Type {
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if alloca, ok := inst.(*Alloca); ok {
				if alloca.Result.ID == regID {
					return alloca.Ty
				}
			}
		}
	}
	return types.IntTySig
}

func copyValueMap(m map[int]Value) map[int]Value {
	cp := make(map[int]Value, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func valueMapsEqual(a, b map[int]Value) bool {
	if len(a) != len(b) {
		return false
	}
	for k, va := range a {
		vb, ok := b[k]
		if !ok || va.String() != vb.String() {
			return false
		}
	}
	return true
}

func phiEntriesEqual(a, b []PhiEntry) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Value.String() != b[i].Value.String() || a[i].Block != b[i].Block {
			return false
		}
	}
	return true
}

func advanceRegisterID(fn *Function, b *builder) {
	maxID := -1
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			switch i := inst.(type) {
			case *Alloca:
				if i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			case *Load:
				if i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			case *BinaryOp:
				if i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			case *Icmp:
				if i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			case *Call:
				if i.Result != nil && i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			case *Gep:
				if i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			case *Bitcast:
				if i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			case *Phi:
				if i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			case *Mov:
				if i.Result.ID > maxID {
					maxID = i.Result.ID
				}
			}
		}
	}
	if maxID >= b.registerID {
		b.registerID = maxID + 1
	}
}
