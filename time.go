package lispy

func Time(li *Lispy) string {
	const htmlstr = `<%s{{if exist "datetime"}} datetime="{{getdel "datetime"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(htmlstr)
}
