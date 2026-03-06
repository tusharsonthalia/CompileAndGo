package ir

// Package ir implements the Out-Of-SSA Translation phase.
// Native ASM execution architectures implicitly cannot process parallel Phi (φ) assignments simultaneously.
// Therefore, the SSA constraint guarantees must be stripped.
//
// Core Algorithm:
// 1. Critical Edge Splitting: Traverses the CFG checking for edges from multiple-successor blocks to multiple-predecessor blocks. It organically splits these with an atomic intermediate block to prevent assignment clobber collisions.
// 2. Phi-Demotion: Disassembles the explicitly formulated φ-nodes into concrete parallel `MOV` instructions appended accurately at the tails of the active predecessor defining nodes.

import "fmt"

func (b *builder) OutOfSSAPass() {
	for _, fn := range b.functions {
		outOfSSA(fn, b)
	}
}

func outOfSSA(fn *Function, b *builder) {
	if len(fn.Blocks) == 0 {
		return
	}

	advanceRegisterID(fn, b)

	splitCriticalEdges(fn, b)

	demotePhis(fn)
}

func splitCriticalEdges(fn *Function, b *builder) {
	newBlocks := make([]*BasicBlock, 0)

	originalBlocks := make([]*BasicBlock, len(fn.Blocks))
	copy(originalBlocks, fn.Blocks)

	for _, A := range originalBlocks {
		if len(A.Successors) <= 1 {
			continue
		}

		for i, B := range A.Successors {
			if len(B.Predecessors) > 1 {

				AB_Label := fmt.Sprintf("%s_%s", A.Label, B.Label)
				A_B := NewBasicBlock(AB_Label)
				b.labelID++

				A_B.AddInstruction(&Branch{Target: B.Label})

				A_B.Predecessors = []*BasicBlock{A}
				A_B.Successors = []*BasicBlock{B}

				A.Successors[i] = A_B

				for j, p := range B.Predecessors {
					if p == A {
						B.Predecessors[j] = A_B
						break
					}
				}

				if term, ok := A.Terminator.(*CondBranch); ok {
					if term.TrueL == B.Label {
						term.TrueL = A_B.Label
					}
					if term.FalseL == B.Label {
						term.FalseL = A_B.Label
					}
				} else if term, ok := A.Terminator.(*Branch); ok {
					if term.Target == B.Label {
						term.Target = A_B.Label
					}
				}

				for _, inst := range B.Instructions {
					if phi, ok := inst.(*Phi); ok {
						for k := range phi.Entries {
							if phi.Entries[k].Block == A.Label {
								phi.Entries[k].Block = A_B.Label
							}
						}
					}
				}

				newBlocks = append(newBlocks, A_B)
			}
		}
	}

	fn.Blocks = append(fn.Blocks, newBlocks...)
}

func demotePhis(fn *Function) {

	blockMap := make(map[string]*BasicBlock)
	for _, block := range fn.Blocks {
		blockMap[block.Label] = block
	}

	for _, block := range fn.Blocks {
		remainingInsts := make([]Instruction, 0)

		for _, inst := range block.Instructions {
			if phi, ok := inst.(*Phi); ok {

				for _, entry := range phi.Entries {
					predBlock := blockMap[entry.Block]
					if predBlock == nil {
						continue
					}

					mov := &Mov{
						Result: phi.Result,
						Src:    entry.Value,
					}

					idx := len(predBlock.Instructions) - 1
					if idx < 0 {
						idx = 0
					}

					newInsts := make([]Instruction, 0, len(predBlock.Instructions)+1)
					newInsts = append(newInsts, predBlock.Instructions[:idx]...)
					newInsts = append(newInsts, mov)
					newInsts = append(newInsts, predBlock.Instructions[idx:]...)
					predBlock.Instructions = newInsts
				}

			} else {
				remainingInsts = append(remainingInsts, inst)
			}
		}

		block.Instructions = remainingInsts
	}
}
