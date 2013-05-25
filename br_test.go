package lispy

import (
	"fmt"
	"testing"
)

func TestBr(t *testing.T) {
	fmt.Println("Br Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(br:)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(br:5)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
