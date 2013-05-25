package lispy

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	fmt.Println("Table Test:\r\n")
	fmt.Println()

	lisp := New()
	code := "(map:(area:rect|coords:0,0,82,126|href:sun.htm|alt:Sun)\r\n"
	code += "(area:circle|coords:90,58,3|href:mercur.htm|alt:Mercury)\r\n"
	code += "(area:circle|coords:124,58,8|href:venus.htm|alt:Venus))\r\n"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
