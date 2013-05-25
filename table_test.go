package lispy

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	fmt.Println("Table Test:\r\n")
	fmt.Println()

	lisp := New()

	code := "(table:(tr:(th:hello)(th:world))(tr:(td:hello)(td:world|colspan:2)))"
	fmt.Println("Input:")
	fmt.Println(code)
	str := lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()

	code = "(table:(tr:(th:hello)(th:world))(tr:(td:hello)(td:world|colspan:2))) (table:(tr:(th:hello)(th:world))(tr:(td:hello)(td:world|colspan:2)))"
	fmt.Println("Input:")
	fmt.Println(code)
	str = lisp.Render(code)
	fmt.Println("Output:")
	fmt.Println(str)
	fmt.Println()
}
