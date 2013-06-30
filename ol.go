package lispy

import (
	"fmt"
)

func Ol(li *Lispy) string {
	str := `<ol`

	if li.Exist("start") {
		str += ` start="` + fmt.Sprint(li.GetInt64Del("start")) + `"`
	}

	if li.Exist("type") {
		str += ` type="` + li.GetDel("type") + `"`
	}

	if li.ExistDel("reverse") {
		str += ` reverse`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</ol>`

	return str
}
