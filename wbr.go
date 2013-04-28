package lispy

import (
	"strings"
)

func Wbr(li *Lispy) string {
	if li.Content == "" {
		return "<wbr>"
	}
	needle := li.Get("needle")
	switch needle {
	case "(", ")", "`":
		return li.Render(li.Content)
	}
	return li.Render(strings.Replace(li.Content, needle, needle+"<wbr>", -1))
}
