package main

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

func TestGetCurrentFunc(t *testing.T) {

	input := `
	package data

	func main() {
		first := make(chan int)
		go sender(first)
		<-first
	}
	`
	var currentFunc string

	f := Parser(input)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			currentFunc = GetCurrentFunc(x)
		}
		return true
	})
	assert.Equal(t, "main", currentFunc, "This should return main")

}

func TestMsgOrigins(t *testing.T) {
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
	var currentFunc string
	var sendOrigin string
	var receiveOrigin string

	f := Parser(input)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			currentFunc = GetCurrentFunc(x)
		case *ast.SendStmt:
			sendOrigin = currentFunc
		case *ast.UnaryExpr:
			receiveOrigin = currentFunc

		}
		return true
	})

	assert.Equal(t, "main", receiveOrigin, "The origin of receive statement is main")
	assert.Equal(t, "sender", sendOrigin, "The origin of send statement is sender function")
}
