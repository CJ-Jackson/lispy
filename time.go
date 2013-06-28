package lispy

import (
	"fmt"
)

func Time(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "datetime"}} datetime="{{.GetDel "datetime"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}
