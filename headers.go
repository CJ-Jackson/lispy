package lispy

import (
	"fmt"
)

func Header(li *Lispy) string {
	const htmlstr = `<%s{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|parse}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}
