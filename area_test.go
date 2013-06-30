package lispy

import (
	"testing"
)

func TestArea(t *testing.T) {
	lisp := New()

	code := "(map:(area:rect|coords:0,0,82,126|href:sun.htm|alt:Sun)"
	code += "(area:circle|coords:90,58,3|href:mercur.htm|alt:Mercury)"
	code += "(area:circle|coords:124,58,8|href:venus.htm|alt:Venus))"

	str := lisp.Render(code)

	if str != `<map><area shape="rect" coords="0,0,82,126" href="sun.htm" alt="Sun"/><area shape="circle" coords="90,58,3" href="mercur.htm" alt="Mercury"/><area shape="circle" coords="124,58,8" href="venus.htm" alt="Venus"/></map>` {
		t.Fail()
	}

	lisp.SetFunc("area", AreaNoFollow)

	str = lisp.Render(code)

	if str != `<map><area shape="rect" coords="0,0,82,126" href="sun.htm" alt="Sun" rel="nofollow"/><area shape="circle" coords="90,58,3" href="mercur.htm" alt="Mercury" rel="nofollow"/><area shape="circle" coords="124,58,8" href="venus.htm" alt="Venus" rel="nofollow"/></map>` {
		t.Fail()
	}
}
