package lispy

func Ol(li *Lispy) string {
	const htmlstr = `<ol{{if exist "start"}} start="{{getint64del "start"}}"{{end}}{{if exist "type"}} type="{{getdel "type"}}"{{end}}{{if existdel "reversed"}} reversed{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</ol>`
	return li.HtmlRender(htmlstr)
}
