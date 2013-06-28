package lispy

func Object(li *Lispy) string {
	const htmlstr = `{{$li := .}}<object{{if .Exist "width"}} width="{{.GetDel "width"}}"{{end}}{{if .Exist "height"}} height="{{.GetDel "height"}}"{{end}}{{if .Exist "data"}} data="{{.GetDel "data"}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</object>`
	return li.HtmlRender(htmlstr)
}

func Param(li *Lispy) string {
	const htmlstr = `{{$li := .}}<param{{if .Exist "name"}} name="{{.GetDel "name"}}"{{end}}{{if .Exist "value"}} value="{{.GetDel "value"}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("value") {
		li.Set("value", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Embed(li *Lispy) string {
	const htmlstr = `{{$li := .}}<embed{{if .Exist "src"}} src="{{.GetDel "src"}}"{{end}}{{if .Exist "type"}} type="{{.GetDel "type"}}"{{end}}{{if .Exist "width"}} width="{{.GetDel "width"}}"{{end}}{{if .Exist "height"}} height="{{.GetDel "height"}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("src") {
		li.Set("src", li.Content)
	}

	return li.HtmlRender(htmlstr)
}
