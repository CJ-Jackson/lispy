package lispy

func Table(li *Lispy) string {
	const htmlstr = `<table{{if existdel "border"}} border="1"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</table>`
	return li.HtmlRender(htmlstr)
}

func TableTd(li *Lispy) string {
	const htmlstr = `<td{{if exist "colspan"}} colspan="{{getdel "colspan"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</td>`
	return li.HtmlRender(htmlstr)
}
