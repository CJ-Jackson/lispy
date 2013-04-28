package lispy

import (
	"fmt"
	"testing"
)

func TestImg(t *testing.T) {
	fmt.Println("Img Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(img`http://example.com/img.png`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(a`(img`http://example.com/img.png`)|href:http://example.com`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
