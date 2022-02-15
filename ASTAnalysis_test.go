package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

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
