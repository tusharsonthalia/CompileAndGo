package ir

// Package ir implements Linear Scan Register Allocation.
// This pass dynamically maps an unbounded amount of SSA virtual-registers to a finite bounded
// pool of physical AArch64 registers (e.g. x19-x28).
//
// Core Algorithm:
// - Maintains a sorted `ActiveList` of live intervals currently held in physical registers.
// - As new intervals arise temporally, expired intervals are popped off the list and their phys-registers returned to the `freePool`.
// - If the `freePool` becomes saturated (0 registers available), the interval with the furthest chronological End date is forcibly Spilled into memory to satisfy the urgent register constraint.

import (
	"fmt"
	"sort"
)

type ActiveList []*LiveInterval

func (a ActiveList) Len() int           { return len(a) }
func (a ActiveList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ActiveList) Less(i, j int) bool { return a[i].End < a[j].End }

var freePool = []string{
	"x0", "x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8",
	"x11", "x12", "x13", "x14", "x15",
	"x19", "x20", "x21", "x22", "x23", "x24", "x25", "x26", "x27", "x28",
}

func allocateRegisters(fn *Function, intervals []*LiveInterval) {
	if len(intervals) == 0 {
		return
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var active ActiveList
	freeRegisters := make([]string, len(freePool))
	copy(freeRegisters, freePool)

	allocationMap := make(map[int]string)

	expireOldIntervals := func(i *LiveInterval) {
		sort.Sort(active)

		retained := make(ActiveList, 0)
		for _, j := range active {
			if j.End >= i.Start {

				retained = append(retained, j)
			} else {

				physReg := allocationMap[j.ID]
				if physReg != "spill" {
					freeRegisters = append(freeRegisters, physReg)
				}
			}
		}
		active = retained
	}

	for _, i := range intervals {
		expireOldIntervals(i)

		if len(active) == len(freePool) {

			spill(i, &active, allocationMap, &freeRegisters)
		} else {

			regToAssign := freeRegisters[0]
			freeRegisters = freeRegisters[1:]

			allocationMap[i.ID] = regToAssign
			active = append(active, i)
			sort.Sort(active)
		}
	}

	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			mutateRegisterAllocations(inst, allocationMap)
		}
	}
}

func spill(i *LiveInterval, active *ActiveList, allocationMap map[int]string, freeRegisters *[]string) {

	spillCandidate := (*active)[len(*active)-1]

	if spillCandidate.End > i.End {

		physReg := allocationMap[spillCandidate.ID]
		allocationMap[spillCandidate.ID] = "spill"
		allocationMap[i.ID] = physReg

		*active = (*active)[:len(*active)-1]
		*active = append(*active, i)
		sort.Sort(active)
	} else {

		allocationMap[i.ID] = "spill"
	}
}

func mutateRegisterAllocations(inst Instruction, allocMap map[int]string) {
	apply := func(v Value) {
		if r, ok := v.(*Register); ok {
			str := allocMap[r.ID]
			if str == "" {
				return
			}

			if str == "spill" {
				r.allocatedPhyReg = "[SPILLED]"
			} else {
				r.allocatedPhyReg = fmt.Sprintf("[%s]", str)
			}
		}
	}

	switch i := inst.(type) {
	case *Load:
		apply(i.Src)
		apply(i.Result)
	case *Store:
		apply(i.Src)
		apply(i.Dst)
	case *BinaryOp:
		apply(i.Left)
		apply(i.Right)
		apply(i.Result)
	case *Icmp:
		apply(i.Left)
		apply(i.Right)
		apply(i.Result)
	case *Call:
		for _, a := range i.Args {
			apply(a)
		}
		if i.Result != nil {
			apply(i.Result)
		}
	case *Gep:
		apply(i.Ptr)
		apply(i.Result)
	case *Bitcast:
		apply(i.Src)
		apply(i.Result)
	case *Mov:
		apply(i.Src)
		apply(i.Result)
	case *Return:
		if i.Val != nil {
			apply(i.Val)
		}
	case *CondBranch:
		apply(i.Cond)
	}
}
