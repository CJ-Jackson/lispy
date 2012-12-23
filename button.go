package lispy

func Button(li *Lispy) string {
	const htmlstr = `<button{{if exist "value"}} value="{{getdel "value"}}"{{end}}{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if exist "type"}} type="{{getdel "type"}}"{{end}}{{if existres "disabled"}} disabled{{end}}{{if existres "autofocus"}} autofocus{{end}}{{if existres "formnovalidate"}} formnovalidate{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`
	return li.HtmlRender(htmlstr)
}
