package lispy

import (
	"fmt"
	"testing"
)

func TestObject(t *testing.T) {
	fmt.Println("Object Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(object`(param`|name:autoplay|value:true`)|data:horse.wav`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(embed`helloworld.swf`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
