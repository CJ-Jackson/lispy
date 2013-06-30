package lispy

import (
	"testing"
)

func TestCommon(t *testing.T) {
	lisp := New()

	code := "(p:Hello World!)"
	str := lisp.Render(code)

	if str != `<p>Hello World!</p>` {
		t.Fail()
	}

	code = "(p:(p:Hello World!))"
	str = lisp.Render(code)

	if str != `<p><p>Hello World!</p></p>` {
		t.Fail()
	}

	code = "(p:(p:Hello World!|class:lisp))"
	str = lisp.Render(code)

	if str != `<p><p class="lisp">Hello World!</p></p>` {
		t.Fail()
	}

	code = "(p:(img:http://example.com/img.png))"
	str = lisp.Render(code)

	if str != `<p><img src="http://example.com/img.png" alt="" title=""/></p>` {
		t.Fail()
	}

	code = "(p:(img:http://example.com/img.png"
	str = lisp.Render(code)

	if str != `<p><img src="http://example.com/img.png" alt="" title=""/></p>` {
		t.Fail()
	}

	code = "(p:(a:(img:http://example.com/img.png)|href:http://example.com))"
	str = lisp.Render(code)

	if str != `<p><a href="http://example.com"><img src="http://example.com/img.png" alt="" title=""/></a></p>` {
		t.Fail()
	}

	code = "(div:Hello World!)"
	str = lisp.Render(code)

	if str != `<div>Hello World!</div>` {
		t.Fail()
	}

	code = "(span:Hello World!)"
	str = lisp.Render(code)

	if str != `<span>Hello World!</span>` {
		t.Fail()
	}

	code = "(span:Hello World!|hello:world)"
	str = lisp.Render(code)

	if str != `<span hello="world">Hello World!</span>` {
		t.Fail()
	}

	code = "(col:Hello World!|span:world)"
	str = lisp.Render(code)

	if str != `<col span="world">Hello World!</col>` {
		t.Fail()
	}

	code = "(blockquote:Hello World!|cite:world)"
	str = lisp.Render(code)

	if str != `<blockquote cite="world">Hello World!</blockquote>` {
		t.Fail()
	}

	code = "(li:Hello World!|value:world)"
	str = lisp.Render(code)

	if str != `<li value="world">Hello World!</li>` {
		t.Fail()
	}

	code = "(blockquote:(blockquote:Hello World!|cite:world)|cite:world)"
	str = lisp.Render(code)

	if str != `<blockquote cite="world"><blockquote cite="world">Hello World!</blockquote></blockquote>` {
		t.Fail()
	}

	code = "(abbr:Hello World!|title:world)"
	str = lisp.Render(code)

	if str != `<abbr title="world">Hello World!</abbr>` {
		t.Fail()
	}

	code = "(p:Test \\|  Test|title:Test ' \\| ' test)"
	str = lisp.Render(code)

	if str != `<p title="Test &#39; | &#39; test">Test |  Test</p>` {
		t.Fail()
	}
}
