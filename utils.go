//Useful functions
package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
)

func GetExpString(a ast.Expr) string {
	return types.ExprString(a)
}

// FindChannelName : Getting all the channels declarations
func FindChannelName(x *ast.AssignStmt) string {
	var left bytes.Buffer
	var right bytes.Buffer
	var chanName string
	//The right side of an assign statement
	for _, val := range x.Rhs {
		right.WriteString(types.ExprString(val))
	}
	rightSideVals := right.String()
	if strings.Contains(rightSideVals, "chan") {
		for _, val := range x.Lhs {
			left.WriteString(types.ExprString(val))
		}
		chanName = left.String()
	}
	right.Reset()
	left.Reset()
	return chanName
}

//CorrelateChans : Correlating the channels by the indices of the Goroutine arguments and functions
//Pars is the parameters of the function
func CorrelateChans(pars []*ast.Field, goruMp map[int]string, chanCor map[string]string) {
	for ind, val := range pars {
		singlePar := val.Names[0]
		if ch, ok := goruMp[ind]; ok {
			//fmt.Println(singlePar.Name, "is at index", ch)
			chanCor[singlePar.Name] = ch
		}
	}
}

// GetChanType : This is used to get the type of channel i.e. Int, String, Float etc
func GetChanType(ch *ast.AssignStmt) string {
	rightSide := ch.Rhs[0].(*ast.CallExpr).Args[0].(*ast.ChanType)
	chanArg := rightSide.Value.(*ast.Ident)

	return chanArg.Name
}

// Testparser : FileSet generation fro the test files
func Testparser(src string) *ast.File {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "test", src, 0)

	return f
}

func GetCurrentFunc(decl *ast.FuncDecl) string {
	return decl.Name.Name
}
func HandleGoroutine(mp map[string]map[int]string, args []ast.Expr, channelMap map[string]bool, name string) {
	mp[name] = make(map[int]string)
	//if the argument is a channel, then add it to the goArgumentsMp with its index
	for i, val := range args {
		valStr := GetExpString(val)
		//Checking if the parameters which are channels
		if _, ok := channelMap[valStr]; ok {
			mp[name][i] = valStr
		}
	}
}
