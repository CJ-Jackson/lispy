package lispy

import (
	"testing"
)

func TestA(t *testing.T) {
	lisp := New()

	code := "(a:http://example.com)"
	str := lisp.Render(code)
	if str != `<a href="http://example.com">http://example.com</a>` {
		t.Fail()
	}

	code = "(a:Hello World!|href:http://example.com)"
	str = lisp.Render(code)
	if str != `<a href="http://example.com">Hello World!</a>` {
		t.Fail()
	}

	code = "(a:Hello World!|href:http://example.com) (a:Hello World!|href:http://example.com)"
	str = lisp.Render(code)
	if str != `<a href="http://example.com">Hello World!</a> <a href="http://example.com">Hello World!</a>` {
		t.Fail()
	}

	code = "(a:Hello World!|href:http://example.com)"

	lisp.SetFunc("a", ANoFollow)

	str = lisp.Render(code)

	if str != `<a href="http://example.com" rel="nofollow">Hello World!</a>` {
		t.Fail()
	}
}
