package lispy

import (
	"fmt"
	"testing"
)

func TestHeaders(t *testing.T) {
	fmt.Println("Header Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(h1`world`) (Hello World)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(h1`(h1`world`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(h1`world|class:test|id:world`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(h1`world|class:(h1`test|id:hello`)|id:(h1`test`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(h1`(h1`test`)|class:(h1`test`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(h1`(h1`test`)|class:(h1`test`)|id:world`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(h1`(h1`world|class:test`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(h1`(h1`wor\r\nld|class:test`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(world) (h1`(h1`world`)`) (world)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(world) (h1`(h1`world`) (world)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(world) (h1`(h1`(h1`world`) (world)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(world) (h1`(h1`world`)(h1`(h1`world`) (world)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(world) (h1`(h1`world`(h1`(h1`world`)`)`)`) (world)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(world) (h1`(world)`) (world)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
