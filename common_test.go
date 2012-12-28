package lispy

import (
	"fmt"
	"testing"
)

func TestCommon(t *testing.T) {
	fmt.Println("Common Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(p`Hello World!`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(p`(p`Hello World!`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(p`(p`Hello World!|class:lisp`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(p`(img`http://example.com/img.png`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(p`(img`http://example.com/img.png"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(p`(a`(img`http://example.com/img.png`)|href:http://example.com`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(div`Hello World!`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(span`Hello World!`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(span`Hello World!|hello:world`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(col`Hello World!|span:world`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(blockquote`Hello World!|cite:world`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(li`Hello World!|value:world`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(blockquote`(blockquote`Hello World!|cite:world`)|cite:world`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(abbr`Hello World!|title:world`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(p`Test \\|  Test|title:Test ' \\| ' test`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Parse(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
