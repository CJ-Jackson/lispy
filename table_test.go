package lispy

import (
	"testing"
)

func TestTable(t *testing.T) {
	lisp := New()

	code := "(table:(tr:(th:hello)(th:world))(tr:(td:hello)(td:world|colspan:2)))"

	str := lisp.Render(code)

	if str != `<table><tr><th>hello</th><th>world</th></tr><tr><td>hello</td><td colspan="2">world</td></tr></table>` {
		t.Fail()
	}

	code = "(table:(tr:(th:hello)(th:world))(tr:(td:hello)(td:world|colspan:2))) (table:(tr:(th:hello)(th:world))(tr:(td:hello)(td:world|colspan:2)))"

	str = lisp.Render(code)

	if str != `<table><tr><th>hello</th><th>world</th></tr><tr><td>hello</td><td colspan="2">world</td></tr></table> <table><tr><th>hello</th><th>world</th></tr><tr><td>hello</td><td colspan="2">world</td></tr></table>` {
		t.Fail()
	}
}
