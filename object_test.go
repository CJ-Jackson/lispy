package lispy

import (
	"testing"
)

func TestObject(t *testing.T) {
	lisp := New()
	code := "(object:(param:|name:autoplay|value:true)|data:horse.wav)"
	str := lisp.Render(code)

	if str != `<object data="horse.wav"><param name="autoplay" value="true"/></object>` {
		t.Fail()
	}

	code = "(embed:helloworld.swf)"
	str = lisp.Render(code)

	if str != `<embed src="helloworld.swf"/>` {
		t.Fail()
	}
}
