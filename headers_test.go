package lispy

import (
	"testing"
)

func TestHeaders(t *testing.T) {

	lisp := New()
	code := "(h1:world) (Hello World)"
	str := lisp.Render(code)

	if str != `<h1>world</h1> (Hello World)` {
		t.Fail()
	}

	code = "(h1:(h1:world))"
	str = lisp.Render(code)

	if str != `<h1><h1>world</h1></h1>` {
		t.Fail()
	}

	code = "(h1:world|class:test|id:world)"
	str = lisp.Render(code)

	if str != `<h1 class="test" id="world">world</h1>` {
		t.Fail()
	}

	code = "(h1:world|class:(h1:test|id:hello)|id:(h1:test))"
	str = lisp.Render(code)

	if str != `<h1 class="(h1:test|id:hello)" id="(h1:test)">world</h1>` {
		t.Fail()
	}

	code = "(h1:(h1:test)|class:(h1:test))"
	str = lisp.Render(code)

	if str != `<h1 class="(h1:test)"><h1>test</h1></h1>` {
		t.Fail()
	}

	code = "(h1:(h1:test)|class:(h1:test)|id:world)"
	str = lisp.Render(code)

	if str != `<h1 class="(h1:test)" id="world"><h1>test</h1></h1>` {
		t.Fail()
	}

	code = "(h1:(h1:world|class:test))"
	str = lisp.Render(code)

	if str != `<h1><h1 class="test">world</h1></h1>` {
		t.Fail()
	}

	code = "(world) (h1:(h1:world)) (world)"
	str = lisp.Render(code)

	if str != `(world) <h1><h1>world</h1></h1> (world)` {
		t.Fail()
	}

	code = "(world) (h1:(h1:(h1:world) (world)"
	str = lisp.Render(code)

	if str != `(world) <h1><h1><h1>world</h1> (world)</h1></h1>` {
		t.Fail()
	}

	code = "(world) (h1:(world)) (world)"
	str = lisp.Render(code)

	if str != `(world) <h1>(world)</h1> (world)` {
		t.Fail()
	}
}
