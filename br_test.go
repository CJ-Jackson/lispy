package lispy

import (
	"testing"
)

func TestBr(t *testing.T) {
	lisp := New()

	code := "(br:)"
	str := lisp.Render(code)

	if str != `<br />` {
		t.Fail()
	}

	code = "(br:5)"
	str = lisp.Render(code)

	if str != `<br /><br /><br /><br /><br />` {
		t.Fail()
	}
}
