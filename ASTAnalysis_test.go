package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestTestparser(t *testing.T) {
	testCase := `
	package data

	func main() {
		first := make(chan int)
		go sender(first, "hello")
		<-first
	}
	`
	fset1 := token.NewFileSet()
	expected, _ := parser.ParseFile(fset1, "test", testCase, 0)
	actual := Testparser(testCase)
	assert.Equal(t, expected, actual, "The AST file types should be the same")

}

func TestGetExpString(t *testing.T) {

	testCase := `
	package data

	func main() {
		first := make(chan int)
		go sender(first, "hello")
		<-first
	}
	`
	var actualArr []string
	expectedArr := []string{"first", "\"hello\""}

	f := Testparser(testCase)
	ast.Inspect(f, func(n ast.Node) bool {

		switch x := n.(type) {
		case *ast.GoStmt:
			goArgs := x.Call.Args

			for _, val := range goArgs {
				valStr := GetExpString(val)
				actualArr = append(actualArr, valStr)
			}
		}
		return true
	})

	for i, val := range actualArr {
		assert.Equal(t, expectedArr[i], actualArr[i], fmt.Sprintf("Should have returned %v", val))
	}

}

func TestCorrelateChansOne(t *testing.T) {

	input := `
		package data

		func main() {
			first := make(chan int)
			go sender(first)
			<-first
		}
		func sender(c chan int) {
			c <- 99
		}
	`

	var currentFunc string = "sender"
	goArgumentsMp := make(map[string]map[int]string)
	chanCorrelation := make(map[string]string)

	goArgumentsMp[currentFunc] = make(map[int]string)
	goArgumentsMp[currentFunc][0] = "first"

	f := Testparser(input)

	ast.Inspect(f, func(n ast.Node) bool {

		switch x := n.(type) {
		case *ast.FuncDecl:
			a := x.Type
			paraVals := a.Params.List
			goMap := goArgumentsMp[currentFunc]
			CorrelateChans(paraVals, goMap, chanCorrelation)
		}
		return true
	})
	errorMsg := "The argument c should correspond to channel first"
	errorMsg2 := "The argument name should not be the same name as channel name"

	assert.Equal(t, "first", chanCorrelation["c"], errorMsg)
	assert.NotEqual(t, "c", chanCorrelation["c"], errorMsg2)
}

func TestCorrelateChansMany(t *testing.T) {
	input := `
		package data

		func main() {
			first := make(chan int)
			second := make(chan float64)
			go sender(first, "amin", second)
		}
		func sender(c chan int, name string, k chan float64) {
		}
	`
	var currentFunc string = "sender"
	goArgumentsMp := make(map[string]map[int]string)
	chanCorrelation := make(map[string]string)

	goArgumentsMp[currentFunc] = make(map[int]string)
	goArgumentsMp[currentFunc][0] = "first"
	goArgumentsMp[currentFunc][2] = "second"

	f := Testparser(input)

	ast.Inspect(f, func(n ast.Node) bool {

		switch x := n.(type) {
		case *ast.FuncDecl:
			a := x.Type
			paraVals := a.Params.List
			goMap := goArgumentsMp[currentFunc]
			CorrelateChans(paraVals, goMap, chanCorrelation)
		}
		return true
	})
	errorMsg := "The argument c should correspond to channel first"

	assert.Equal(t, "first", chanCorrelation["c"], errorMsg)
	assert.Equal(t, "second", chanCorrelation["k"], errorMsg)

}

func TestHandleGoroutine(t *testing.T) {
	input := `
	package data

	func main() {
		first := make(chan int)
		go sender(first)
		<-first
	}
	func sender(c chan int) {
		c <- 99
	}
	`
	//Mocking data for testing
	goName := "sender"
	goArgumentsMp := make(map[string]map[int]string)
	chanMap := make(map[string]bool)
	chanMap["first"] = true

	f := Testparser(input)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GoStmt:
			st := x.Call.Fun.(*ast.Ident)
			goArgs := x.Call.Args
			argName := st.Name
			HandleGoroutine(goArgumentsMp, goArgs, chanMap, argName)
		}
		return true
	})
	expected := "first"
	actual := goArgumentsMp[goName][0]

	assert.Equal(t, expected, actual, "The channel named first should be at index 0")

}
