package lispy

import (
	"testing"
)

func TestImg(t *testing.T) {
	lisp := New()
	code := "(img:http://example.com/img.png)"
	str := lisp.Render(code)

	if str != `<img src="http://example.com/img.png" alt="" title=""/>` {
		t.Fail()
	}

	code = "(a:(img:http://example.com/img.png)|href:http://example.com)"
	str = lisp.Render(code)

	if str != `<a href="http://example.com"><img src="http://example.com/img.png" alt="" title=""/></a>` {
		t.Fail()
	}
}
