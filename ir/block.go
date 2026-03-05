package ir

import (
	"fmt"
	"strings"
)

type BasicBlock struct {
	Label        string
	Instructions []Instruction
	Terminator   Instruction
	Predecessors []*BasicBlock
	Successors   []*BasicBlock
}

func NewBasicBlock(label string) *BasicBlock {
	return &BasicBlock{
		Label:        label,
		Instructions: make([]Instruction, 0),
	}
}

func (b *BasicBlock) AddInstruction(inst Instruction) {
	if b.Terminator != nil {
		return
	}
	b.Instructions = append(b.Instructions, inst)

	switch inst.(type) {
	case *Return, *Branch, *CondBranch:
		b.Terminator = inst
	}
}

func (b *BasicBlock) String() string {
	var out strings.Builder
	out.WriteString(fmt.Sprintf("%s:\n", b.Label))
	for _, inst := range b.Instructions {
		out.WriteString(fmt.Sprintf("  %s\n", inst.String()))
	}
	return out.String()
}
