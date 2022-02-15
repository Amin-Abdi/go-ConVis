package main

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

//Test opposite cases i.e, isEven = true and isEven = false

func TestGetChannelName(t *testing.T) {
	testcase := `
	package main

	func main() {
		first := make(chan int)
	}
	`
	f := Testparser(testcase)
	var expected string

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			expected = FindChannelName(x)
		}
		return true
	})
	assert.Equal(t, "first", expected, "Should be Equal")
	assert.NotEqual(t, "second", expected, "Should be Equal")
}

func TestNegativeChannelName(t *testing.T) {
	testcase := `
	package main

	func main() {
		value := "Hello World"
	}
	`
	f := Testparser(testcase)
	var expected string

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			expected = FindChannelName(x)
		}
		return true
	})

	assert.Equal(t, expected, "", "Should return an empty string")
}
