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
		height := 22.01
		first := make(chan int)
	}
	`
	f := Parser(testcase)
	var expected string

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			expected = FindChannelName(x)
		}
		return true
	})
	assert.Equal(t, "first", expected, "Should be Equal")
	assert.NotEqual(t, "height", expected, "Should be Equal")
}

func TestNegativeChannelName(t *testing.T) {
	testcase := `
	package main

	func main() {
		value := "Hello World"
	}
	`
	f := Parser(testcase)
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

func TestChannelTypes(t *testing.T) {
	input1 := `
	package main

	func main() {
		integerChan := make(chan int)
	}
	`
	input2 := `
	package main

	func main() {
		strChan := make(chan string)
	}
	`
	input3 := `
	package main

	func main() {
		floatChan1 := make(chan float32)
	}
	`
	input4 := `
	package main

	func main() {
		floatChan2 := make(chan float64)
	}
	`

	//Preparing the mock data for the input
	names := []string{"integerChan", "strChan", "floatChan1", "floatChan2"}
	testSlice := []string{input1, input2, input3, input4}
	//Channel used for checking
	chanTypeMap := make(map[string]string)

	for i := 0; i < len(testSlice); i++ {
		f := Parser(testSlice[i])

		chanName := names[i]

		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.AssignStmt:
				chanTypeMap[chanName] = GetChanType(x)
			}
			return true
		})
	}

	intActual := chanTypeMap[names[0]]
	strActual := chanTypeMap[names[1]]
	float1Actual := chanTypeMap[names[2]]
	float2Actual := chanTypeMap[names[3]]

	assert.Equal(t, "int", intActual, "Should return an int value")
	assert.Equal(t, "string", strActual, "Should return a string value")
	assert.Equal(t, "float32", float1Actual, "Should return a float32 value")
	assert.Equal(t, "float64", float2Actual, "Should return a float64 value")

}
