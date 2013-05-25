package lispy

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println("A Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(a:http://example.com)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	lisp = New()
	code = "(a:Hello World!|href:http://example.com)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	lisp = New()
	code = "(a:Hello World!|href:http://example.com) (a:Hello World!|href:http://example.com)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
