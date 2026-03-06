package ir

import "fmt"

// LinearScanPass drives the full register allocation pipeline per function:
// linearize blocks in RPO, compute liveness, build intervals, then allocate.
func (b *builder) LinearScanPass() {
	for _, fn := range b.functions {

		rpoBlocks, instPosMap := linearizeCodeRPO(fn)
		if len(rpoBlocks) == 0 {
			continue
		}

		_, liveOut := buildBlockLiveness(rpoBlocks)

		intervals := buildLiveIntervals(rpoBlocks, instPosMap, liveOut)

		allocateRegisters(fn, intervals)
	}
}

// linearizeCodeRPO orders blocks in reverse post-order, which gives us a sensible
// linear layout for the liveness analysis (dominators come before dominated blocks).
// Also assigns a sequential position number to each instruction.
func linearizeCodeRPO(fn *Function) ([]*BasicBlock, map[Instruction]int) {
	if len(fn.Blocks) == 0 {
		return nil, nil
	}

	visited := make(map[string]bool)
	postOrder := make([]*BasicBlock, 0)

	var dfs func(b *BasicBlock)
	dfs = func(curr *BasicBlock) {
		visited[curr.Label] = true
		for _, succ := range curr.Successors {
			if !visited[succ.Label] {
				dfs(succ)
			}
		}
		postOrder = append(postOrder, curr)
	}

	dfs(fn.Blocks[0])

	// Reverse post-order = reverse of the postOrder list
	rpoBlocks := make([]*BasicBlock, len(postOrder))
	for i, j := 0, len(postOrder)-1; j >= 0; i, j = i+1, j-1 {
		rpoBlocks[i] = postOrder[j]
	}

	// Reorder function blocks to match RPO (important for the backend too)
	fn.Blocks = rpoBlocks

	// Assign each instruction a position for interval computation
	instPosMap := make(map[Instruction]int)
	pos := 0

	for _, block := range rpoBlocks {
		for _, inst := range block.Instructions {
			instPosMap[inst] = pos
			pos++
		}
	}

	return rpoBlocks, instPosMap
}

// buildBlockLiveness runs the iterative backward dataflow analysis to compute
// LiveIn/LiveOut sets for each block. Classic equations:
//
//	LiveOut[B] = union of LiveIn[S] for all successors S
//	LiveIn[B]  = Use[B] union (LiveOut[B] - Def[B])
//
// Iterates until reaching a fixed point.
func buildBlockLiveness(blocks []*BasicBlock) (map[string]map[int]bool, map[string]map[int]bool) {

	useMap := make(map[string]map[int]bool)
	defMap := make(map[string]map[int]bool)
	liveIn := make(map[string]map[int]bool)
	liveOut := make(map[string]map[int]bool)

	for _, b := range blocks {
		useMap[b.Label] = make(map[int]bool)
		defMap[b.Label] = make(map[int]bool)
		liveIn[b.Label] = make(map[int]bool)
		liveOut[b.Label] = make(map[int]bool)

		for _, inst := range b.Instructions {
			parseInstUsesAndDefs(inst, useMap[b.Label], defMap[b.Label])
		}
	}

	// Fixed-point iteration, bottom-up
	changed := true
	for changed {
		changed = false

		for i := len(blocks) - 1; i >= 0; i-- {
			bl := blocks[i]
			l := bl.Label

			oldOut := fmt.Sprint(liveOut[l])
			oldIn := fmt.Sprint(liveIn[l])

			// LiveOut = union of successor LiveIn sets
			for _, succ := range bl.Successors {
				for r := range liveIn[succ.Label] {
					liveOut[l][r] = true
				}
			}

			// LiveIn = Use union (LiveOut - Def)
			for r := range useMap[l] {
				liveIn[l][r] = true
			}

			for r := range liveOut[l] {
				if !defMap[l][r] {
					liveIn[l][r] = true
				}
			}

			if oldOut != fmt.Sprint(liveOut[l]) || oldIn != fmt.Sprint(liveIn[l]) {
				changed = true
			}
		}
	}

	return liveIn, liveOut
}

// parseInstUsesAndDefs classifies each register operand as a "use" or "def".
// A register is a use if it's read before being defined in this block.
func parseInstUsesAndDefs(inst Instruction, uses, defs map[int]bool) {
	addUse := func(v Value) {
		if reg, ok := v.(*Register); ok {
			// Only count as a use if we haven't already defined it in this block
			if !defs[reg.ID] {
				uses[reg.ID] = true
			}
		}
	}

	addDef := func(v Value) {
		if reg, ok := v.(*Register); ok {

			defs[reg.ID] = true
		}
	}

	switch i := inst.(type) {
	case *Load:
		addUse(i.Src)
		addDef(i.Result)
	case *Store:
		addUse(i.Src)
		addUse(i.Dst)
	case *BinaryOp:
		addUse(i.Left)
		addUse(i.Right)
		addDef(i.Result)
	case *Icmp:
		addUse(i.Left)
		addUse(i.Right)
		addDef(i.Result)
	case *Call:
		for _, arg := range i.Args {
			addUse(arg)
		}
		if i.Result != nil {
			addDef(i.Result)
		}
	case *Gep:
		addUse(i.Ptr)
		addDef(i.Result)
	case *Bitcast:
		addUse(i.Src)
		addDef(i.Result)
	case *Mov:
		addUse(i.Src)
		addDef(i.Result)
	case *Return:
		if i.Val != nil {
			addUse(i.Val)
		}
	case *CondBranch:
		addUse(i.Cond)
	}
}

type LiveInterval struct {
	ID    int
	Start int
	End   int
}

func buildLiveIntervals(blocks []*BasicBlock, posMap map[Instruction]int, liveOut map[string]map[int]bool) []*LiveInterval {
	ivMap := make(map[int]*LiveInterval)

	getMaxPos := func(b *BasicBlock) int {
		if len(b.Instructions) == 0 {
			return 0
		}
		return posMap[b.Instructions[len(b.Instructions)-1]]
	}

	for _, b := range blocks {
		for _, inst := range b.Instructions {
			pos := posMap[inst]

			track := func(v Value, isDef bool) {
				if r, ok := v.(*Register); ok {
					if _, exists := ivMap[r.ID]; !exists {
						ivMap[r.ID] = &LiveInterval{ID: r.ID, Start: pos, End: pos}
					} else {
						if isDef && pos < ivMap[r.ID].Start {
							ivMap[r.ID].Start = pos
						}
						if pos > ivMap[r.ID].End {
							ivMap[r.ID].End = pos
						}
					}
				}
			}

			switch i := inst.(type) {
			case *Load:
				track(i.Src, false)
				track(i.Result, true)
			case *Store:
				track(i.Src, false)
				track(i.Dst, false)
			case *BinaryOp:
				track(i.Left, false)
				track(i.Right, false)
				track(i.Result, true)
			case *Icmp:
				track(i.Left, false)
				track(i.Right, false)
				track(i.Result, true)
			case *Call:
				for _, a := range i.Args {
					track(a, false)
				}
				if i.Result != nil {
					track(i.Result, true)
				}
			case *Gep:
				track(i.Ptr, false)
				track(i.Result, true)
			case *Bitcast:
				track(i.Src, false)
				track(i.Result, true)
			case *Mov:
				track(i.Src, false)
				track(i.Result, true)
			case *Return:
				if i.Val != nil {
					track(i.Val, false)
				}
			case *CondBranch:
				track(i.Cond, false)
			}
		}
	}

	// Extend intervals for registers that are live across block boundaries
	for _, b := range blocks {
		boundaryPos := getMaxPos(b)
		for id := range liveOut[b.Label] {
			if bnd, ok := ivMap[id]; ok {
				if boundaryPos > bnd.End {
					bnd.End = boundaryPos
				}
			}
		}
	}

	results := make([]*LiveInterval, 0, len(ivMap))
	for _, iv := range ivMap {
		results = append(results, iv)
	}
	return results
}
