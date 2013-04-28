package lispy

import (
	"fmt"
	"testing"
)

func TestLispCode(t *testing.T) {
	fmt.Println("Output Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(hello`world`) (Hello World)"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "world (Hello World)" {
		t.Fail()
	}

	code = "(hello``)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "" {
		t.Fail()
	}

	code = "(hello`(hello`world`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "(hello`world`)" {
		t.Fail()
	}

	code = "(hello`(hello`(hello`world`)`)`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "(hello`(hello`world`)`)" {
		t.Fail()
	}

	code = "(hello`world|key:value`)"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "world|key:value" {
		t.Fail()
	}

	code = "(hello`world|key:value"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "world|key:value" {
		t.Fail()
	}

	code = "(hello`(hello`world|key:value"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "(hello`world|key:value" {
		t.Fail()
	}

	code = "<test />"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "&lt;test /&gt;" {
		t.Fail()
	}

	lisp.DisableAutoEscape()

	str = lisp.Render(code)

	if str != "<test />" {
		t.Fail()
	}

	lisp = New()
	lisp.EnableAutoLineBreak()
	code = "hello\r\n\r\n"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	if str != "hello<br /><br />" {
		t.Fail()
	}
}
