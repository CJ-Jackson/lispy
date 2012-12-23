package lispy

import (
	"fmt"
	"testing"
)

func TestWbr(t *testing.T) {
	fmt.Println("Wbr Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(wbr``)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
