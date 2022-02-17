package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	//Getting the source file
	srcPath := os.Args[1]
	astAnalysis(srcPath)

	startServer()
}

type Representation struct {
	TypeOp      string `json:"operation"`
	Name        string `json:"name"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Value       string `json:"value"`
	Condition   bool   `json:"condition"`
}

// Operations : list of representations
var Operations []Representation

func astAnalysis(source string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, source, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	traverseAST(f)
	//ast.Print(fset, f)

}

func traverseAST(f *ast.File) {
	var currentFunc string // The name of the current function

	//This map is used to check if a variable is a channel or not
	//To be used with the goArgumentsMp
	channelMap := make(map[string]bool)
	//This is for the channel types
	chanTypeMap := make(map[string]string)
	/*
		This is for storing the goroutine arguments and the values
		it stores a map of key: "GoRoutine name" and values are the channel names and their indices
	*/
	goArgumentsMp := make(map[string]map[int]string)
	//This is for the channel Correlation
	chanCorrelation := make(map[string]string)

	//Representation{TypeOp: "goroutine", Origin: currentFunc, Name: st.Name}
	Operations = append(Operations, Representation{TypeOp: "goroutine", Name: "main"})

	//Check for if statements
	var conditional bool
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			chanName := FindChannelName(x)

			//If a channel is found then add to the channel map
			if chanName != "" {
				channelMap[chanName] = true
				//fmt.Println(chanName, "uses", GetChanType(x))
				chanTypeMap[chanName] = GetChanType(x)
			}

		case *ast.IfStmt:
			conditional = true

		case *ast.SelectStmt:
			conditional = true

		case *ast.FuncDecl:
			conditional = false
			currentFunc = GetCurrentFunc(x)
			a := x.Type
			paraVals := a.Params.List
			goMap := goArgumentsMp[currentFunc]

			CorrelateChans(paraVals, goMap, chanCorrelation)

		case *ast.GoStmt:
			st := x.Call.Fun.(*ast.Ident)
			goArgs := x.Call.Args
			currGo := Representation{TypeOp: "goroutine", Name: st.Name, Origin: currentFunc}
			argName := st.Name
			HandleGoroutine(goArgumentsMp, goArgs, channelMap, argName)
			//The currentFunc can be used to get the origin of the goroutine
			Operations = append(Operations, currGo)

		case *ast.UnaryExpr:
			recName := x.X.(*ast.Ident)
			recStmt := recName.Name

			//Checking for the channel name inside the channel Correlation map
			if val, ok := chanCorrelation[recName.Name]; ok {
				recStmt = val
			}
			//This is to avoid other types of unary expressions from being recorded
			//Check if the origin of the receive stmt is a channel
			if _, ok := channelMap[recStmt]; ok {
				myRec := Representation{TypeOp: "receive", Origin: recStmt, Destination: currentFunc, Condition: conditional}
				Operations = append(Operations, myRec)
			}
		case *ast.SendStmt:
			dest := x.Chan.(*ast.Ident).Name
			//Check if the sendStmt Channel name is in the correlation, if not then just use that chan name
			if val, ok := chanCorrelation[dest]; ok {
				dest = val
			}
			valSent := strings.ToUpper(chanTypeMap[dest])
			mySend := Representation{TypeOp: "send", Origin: currentFunc, Destination: dest, Value: valSent, Condition: conditional}
			Operations = append(Operations, mySend)
		}
		return true
	})

	fmt.Println("===============================")

	for _, val := range Operations {
		fmt.Println(val)
	}
}
