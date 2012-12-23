package lispy

func Object(li *Lispy) string {
	const htmlstr = `<object{{if exist "width"}} width="{{getdel "width"}}"{{end}}{{if exist "height"}} height="{{getdel "height"}}"{{end}}{{if exist "data"}} data="{{getdel "data"}}"{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</object>`
	return li.HtmlRender(htmlstr)
}

func Param(li *Lispy) string {
	const htmlstr = `<param{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if exist "value"}} value="{{getdel "value"}}"{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("value") {
		li.Set("value", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Embed(li *Lispy) string {
	const htmlstr = `<embed{{if exist "src"}} src="{{getdel "src"}}"{{end}}{{if exist "type"}} type="{{getdel "type"}}"{{end}}{{if exist "width"}} width="{{getdel "width"}}"{{end}}{{if exist "height"}} height="{{getdel "height"}}"{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("src") {
		li.Set("src", li.Content)
	}

	return li.HtmlRender(htmlstr)
}
