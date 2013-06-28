package lispy

func Form(li *Lispy) string {
	const htmlstr = `{{$li := .}}<form{{if .Exist "action"}} action="{{.GetDel "action"|html}}"{{end}}{{if .Exist "method"}} method="{{.GetDel "method"|html}}"{{end}}{{if .ExistDel "novalidate"}} novalidate{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("action") {
		li.Set("action", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Input(li *Lispy) string {
	const htmlstr = `{{$li := .}}<input{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{if .Exist "value"}} value="{{.GetDel "value"|html}}"{{end}}{{if .ExistDel "checked"}} checked{{end}}{{if .ExistRes "disabled"}} disabled{{end}}{{if .ExistDel "autofocus"}} autofocus{{end}}{{if .ExistDel "formnovalidate"}} formnovalidate{{end}}{{if .ExistDel "multiple"}} multiple{{end}}{{if .ExistDel "readonly"}} readonly{{end}}{{if .ExistDel "required"}} required{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Textarea(li *Lispy) string {
	const htmlstr = `{{$li := .}}<textarea{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{if .Exist "value"}} value="{{.GetDel "value"|html}}"{{end}}{{if .ExistRes "disabled"}} disabled{{end}}{{if .ExistDel "autofocus"}} autofocus{{end}}{{if .ExistDel "readonly"}} readonly{{end}}{{if .ExistDel "required"}} required{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</textarea>`

	return li.HtmlRender(htmlstr)
}

func Select(li *Lispy) string {
	const htmlstr = `{{$li := .}}<select{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{if .ExistRes "disabled"}} disabled{{end}}{{if .ExistDel "autofocus"}} autofocus{{end}}{{if .ExistDel "multiple"}} multiple{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</select>`

	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Option(li *Lispy) string {
	const htmlstr = `{{$li := .}}<option{{if .Exist "value"}} value="{{.GetDel "value"|html}}"{{end}}{{if .ExistRes "disabled"}} disabled{{end}}{{if .ExistDel "selected"}} selected{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</option>`

	return li.HtmlRender(htmlstr)
}

func Output(li *Lispy) string {
	const htmlstr = `{{$li := .}}<output{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{if .Exist "for"}} for="{{.GetDel "for"|html}}"{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}}></output>`

	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Label(li *Lispy) string {
	const htmlstr = `{{$li := .}}<label{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{if .Exist "for"}} for="{{.GetDel "for"|html}}"{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</label>`
	return li.HtmlRender(htmlstr)
}

func FieldSet(li *Lispy) string {
	const htmlstr = `{{$li := .}}<fieldset{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{if .ExistRes "disabled"}} disabled{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</fieldset>`
	return li.HtmlRender(htmlstr)
}

func Keygen(li *Lispy) string {
	const htmlstr = `{{$li := .}}<keygen{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{if .Exist "keytype"}} keytype="{{.GetDel "keytype"|html}}"{{end}}{{if .ExistRes "disabled"}} disabled{{end}}{{if .ExistRes "challenge"}} challenge{{end}}{{if .ExistDel "autofocus"}} autofocus{{end}}{{range .GetNames}} {{.}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	return li.HtmlRender(htmlstr)
}
