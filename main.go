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

}

type Representation struct {
	TypeOp      string `json:"operation"`
	Name        string `json:"name"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Value       string `json:"value"`
}

// Operations : list of representations
var Operations []Representation

func astAnalysis(source string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, source, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	//ast.Print(fset, f)

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

	//Send stmt tests
	//var mySends string[]

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

		case *ast.FuncDecl:
			currentFunc = x.Name.Name

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

			currGo := Representation{TypeOp: "goroutine", Name: st.Name, Origin: currentFunc}
			//currGo := Creation{TypeOp: "creation", Name: st.Name, Parent: currentFunc}

			//if the argument is a channel, then add it to the goArgumentsMp with its index
			for i, val := range goArgs {
				valStr := GetExpString(val)
				//Checking if the parameters which are channels
				if _, ok := channelMap[valStr]; ok {
					goArgumentsMp[st.Name][i] = valStr
				}
			}
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
				myRec := Representation{TypeOp: "receive", Origin: recStmt, Destination: currentFunc}
				Operations = append(Operations, myRec)
			}

		//fmt.Printf("The receive statement is from channel %v to the Goroutine \"%v\" \n", recStmt, currentFunc)
		case *ast.SendStmt:
			//valSent := x.Value.(*ast.BasicLit).Kind.String()
			//The channel where the value is being sent
			dest := x.Chan.(*ast.Ident).Name

			//Check if the sendStmt Channel name is in the correlation, if not then just use that chan name
			if val, ok := chanCorrelation[dest]; ok {
				dest = val
			}

			valSent := strings.ToUpper(chanTypeMap[dest])

			//mySend := SendRec{TypeOp: "send", Origin: currentFunc, Destination: dest, Value: valSent}
			mySend := Representation{TypeOp: "send", Origin: currentFunc, Destination: dest, Value: valSent}
			Operations = append(Operations, mySend)
			//fmt.Printf("The value %v is sent to the channel %v from Goroutine \"%v\" \n", valSent, dest, currentFunc)
			//fmt.Println("Sending to channel:", dest, fset.Position(x.Pos()), "a value of", valSent)
		}
		return true
	})

	fmt.Println("===============================")
	//fmt.Println("channels:", channelMap)
	//fmt.Println("Channel correlation:", chanCorrelation)
	//fmt.Println("GoArgumentMap:", goArgumentsMp)
	//fmt.Println("The Representation List:\n", Operations)

	for _, val := range Operations {
		fmt.Println(val)
	}
	HandleRequests()
}
