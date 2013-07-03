package lispy

import (
	"testing"
)

func TestLispCode(t *testing.T) {
	lisp := New()
	code := "(hello:world) (Hello World)"
	str := lisp.Render(code)

	if str != "world (Hello World)" {
		t.Fail()
	}

	code = "(hello:)"
	str = lisp.Render(code)

	if str != "" {
		t.Fail()
	}

	code = "(hello:(hello:world))"
	str = lisp.Render(code)

	if str != "(hello:world)" {
		t.Fail()
	}

	code = "(hello:(hello:(hello:world)))"
	str = lisp.Render(code)

	if str != "(hello:(hello:world))" {
		t.Fail()
	}

	code = "(hello:world|key:value)"
	str = lisp.Render(code)

	if str != "world|key:value" {
		t.Fail()
	}

	code = "(hello:world|key:value"
	str = lisp.Render(code)

	if str != "world|key:value" {
		t.Fail()
	}

	code = "(hello:(hello:world|key:value"
	str = lisp.Render(code)

	if str != "(hello:world|key:value" {
		t.Fail()
	}

	code = "<test />"
	str = lisp.Render(code)

	if str != "&lt;test /&gt;" {
		t.Fail()
	}
}
