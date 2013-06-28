package lispy

func Table(li *Lispy) string {
	const htmlstr = `{{$li := .}}<table{{if .ExistDel "border"}} border="1"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</table>`
	return li.HtmlRender(htmlstr)
}

func TableTd(li *Lispy) string {
	const htmlstr = `{{$li := .}}<td{{if .Exist "colspan"}} colspan="{{.GetDel "colspan"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</td>`
	return li.HtmlRender(htmlstr)
}
