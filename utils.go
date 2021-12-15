//Useful functions
package main

import (
	"bytes"
	"go/ast"
	"go/types"
	"strings"
)

func GetExpString(a ast.Expr) string {
	return types.ExprString(a)
}

func FindChannelName(x *ast.AssignStmt) string {
	var left bytes.Buffer
	var right bytes.Buffer
	var chanName string

	//The right side of an assign statement
	for _, val := range x.Rhs {
		right.WriteString(types.ExprString(val))
	}
	rightSideVals := right.String()
	//fmt.Println("RightSide:", rightSideVals)

	if strings.Contains(rightSideVals, "chan") {
		for _, val := range x.Lhs {
			left.WriteString(types.ExprString(val))
		}
		//fmt.Println("channel Name:", left.String())
		chanName = left.String()
	}

	right.Reset()
	left.Reset()

	return chanName
}
