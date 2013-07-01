package lispy

import (
	"testing"
)

func TestWbr(t *testing.T) {
	lisp := New()
	code := "(wbr:)"

	str := lisp.Render(code)

	if str != `<wbr>` {
		t.Fail()
	}

	code = "(wbr:hello-world|needle:-)"

	str = lisp.Render(code)

	if str != `hello-<wbr>world` {
		t.Fail()
	}
}
