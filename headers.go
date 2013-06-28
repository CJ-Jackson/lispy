package lispy

import (
	"fmt"
)

func Header(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}
