package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	//Getting the source file
	srcPath := os.Args[1]
	astAnalysis(srcPath)

	//fmt.Println("The goroutines are:", myGoroutines)
}

func astAnalysis(source string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, source, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	//ast.Print(fset, f)

	var funcs []string
	var currentFunc string

	//This map is used to check if a variable is a channel or not
	channelMap := make(map[string]bool)

	/*
		This is for storing the goroutine arguments and the values
		it stores a map of key: "GoRoutine name" and values are the channel names and their indices
	*/
	goArgumentsMp := make(map[string]map[int]string)

	//This is for the channel Correlation
	chanCorrelation := make(map[string]string)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			chanName := FindChannelName(x)
			//fmt.Println("Found channel:", chanName)

			//If a channel is found then add to the channel map
			if chanName != "" {
				channelMap[chanName] = true
			}

		case *ast.FuncDecl:
			currentFunc = x.Name.Name
			funcs = append(funcs, currentFunc)
			//Getting the *ast.FuncType to access the parameters
			a := x.Type
			//Getting the parameters of the function
			paraVals := a.Params.List
			//Returning the map that is stored by the func name i.e. sender
			goMap := goArgumentsMp[currentFunc]
			CorrelateChans(paraVals, goMap, chanCorrelation)

		case *ast.GoStmt:
			st := x.Call.Fun.(*ast.Ident)
			goArgs := x.Call.Args
			goArgumentsMp[st.Name] = make(map[int]string)

			for i, val := range goArgs {
				valStr := GetExpString(val)
				//Checking if the parameters which are channels
				if _, ok := channelMap[valStr]; ok {
					goArgumentsMp[st.Name][i] = valStr
				}
			}

		case *ast.SendStmt:
			valSent := x.Value.(*ast.BasicLit).Kind.String()
			//The channel where the value is being sent
			dest := x.Chan.(*ast.Ident).Name
			fmt.Printf("The value %v is sent to the channel %v from Goroutine \"%v\" \n", valSent, chanCorrelation[dest], currentFunc)

		case *ast.UnaryExpr:
			recName := x.X.(*ast.Ident)
			fmt.Printf("The receieve statement is from channel %v to the Goroutine \"%v\" \n", recName.Name, currentFunc)
		}
		return true
	})

	fmt.Println("===============")
	//fmt.Println("The functions are:", funcs)
	fmt.Println("channels:", channelMap)
	fmt.Println("Channel correlation:", chanCorrelation)

}

func CorrelateChans(pars []*ast.Field, goruMp map[int]string, chanCor map[string]string) {

	for ind, val := range pars {
		singlePar := val.Names[0]
		if ch, ok := goruMp[ind]; ok {
			//fmt.Println(singlePar.Name, "is at index", index)
			chanCor[singlePar.Name] = ch
		}
	}

}

func GetFuncParams(a *ast.FuncType) {
	paraVals := a.Params.List
	for index, val := range paraVals {
		singlePar := val.Names[0]
		fmt.Println(singlePar.Name, "is at index", index)
	}
}
