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
		return li.Parse(li.Content)
	}
	return li.Parse(strings.Replace(li.Content, needle, needle+"<wbr>", -1))
}
