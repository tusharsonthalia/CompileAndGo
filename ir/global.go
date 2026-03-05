package ir

import (
	"fmt"
	"golite/types"
)

type Global struct {
	Name      string
	Ty        types.Type
	Value     string
	IsConst   bool
	IsPrivate bool
}

func (g *Global) String() string {
	scope := "global"
	if g.IsConst {
		scope = "constant"
	}
	visibility := ""
	if g.IsPrivate {
		visibility = "private unnamed_addr "
	}

	val := g.Value
	if val == "" {
		val = "zeroinitializer"
	}

	return fmt.Sprintf("@%s = %s%s %s %s", g.Name, visibility, scope, LLVMType(g.Ty), val)
}
