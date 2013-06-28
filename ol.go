package lispy

func Ol(li *Lispy) string {
	const htmlstr = `{{$li := .}}<ol{{if .Exist "start"}} start="{{.GetInt64Del "start"|html}}"{{end}}{{if .Exist "type"}} type="{{.GetDel "type"|html}}"{{end}}{{if .ExistDel "reversed"}} reversed{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</ol>`
	return li.HtmlRender(htmlstr)
}
