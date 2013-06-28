package lispy

func Button(li *Lispy) string {
	const htmlstr = `{{$li := .}}<button{{if .Exist "value"}} value="{{.GetDel "value"|html}}"{{end}}{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{if .Exist "type"}} type="{{.GetDel "type"|html}}"{{end}}{{if ExistRes "disabled"}} disabled{{end}}{{if ExistRes "autofocus"}} autofocus{{end}}{{if ExistRes "formnovalidate"}} formnovalidate{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}} />`
	return li.HtmlRender(htmlstr)
}
