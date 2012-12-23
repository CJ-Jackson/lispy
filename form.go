package lispy

func Form(li *Lispy) string {
	const htmlstr = `<form{{if exist "action"}} action="{{getdel "action"}}"{{end}}{{if exist "method"}} method="{{getdel "method"}}"{{end}}{{if existdel "novalidate"}} novalidate{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("action") {
		li.Set("action", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Input(li *Lispy) string {
	const htmlstr = `<input{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if exist "value"}} value="{{getdel "value"}}"{{end}}{{if existdel "checked"}} checked{{end}}{{if existres "disabled"}} disabled{{end}}{{if existdel "autofocus"}} autofocus{{end}}{{if existdel "formnovalidate"}} formnovalidate{{end}}{{if existdel "multiple"}} multiple{{end}}{{if existdel "readonly"}} readonly{{end}}{{if existdel "required"}} required{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Textarea(li *Lispy) string {
	const htmlstr = `<textarea{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if exist "value"}} value="{{getdel "value"}}"{{end}}{{if existres "disabled"}} disabled{{end}}{{if existdel "autofocus"}} autofocus{{end}}{{if existdel "readonly"}} readonly{{end}}{{if existdel "required"}} required{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</textarea>`

	return li.HtmlRender(htmlstr)
}

func Select(li *Lispy) string {
	const htmlstr = `<select{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if existres "disabled"}} disabled{{end}}{{if existdel "autofocus"}} autofocus{{end}}{{if existdel "multiple"}} multiple{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</select>`

	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Option(li *Lispy) string {
	const htmlstr = `<option{{if exist "value"}} value="{{getdel "value"}}"{{end}}{{if existres "disabled"}} disabled{{end}}{{if existdel "selected"}} selected{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</option>`

	return li.HtmlRender(htmlstr)
}

func Output(li *Lispy) string {
	const htmlstr = `<output{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if exist "for"}} for="{{getdel "for"}}"{{end}}{{range names}} {{.}}="{{get .}}"{{end}}></output>`

	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Label(li *Lispy) string {
	const htmlstr = `<label{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if exist "for"}} for="{{getdel "for"}}"{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</label>`
	return li.HtmlRender(htmlstr)
}

func FieldSet(li *Lispy) string {
	const htmlstr = `<fieldset{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if existres "disabled"}} disabled{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</fieldset>`
	return li.HtmlRender(htmlstr)
}

func Keygen(li *Lispy) string {
	const htmlstr = `<keygen{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{if exist "keytype"}} keytype="{{getdel "keytype"}}"{{end}}{{if existres "disabled"}} disabled{{end}}{{if existres "challenge"}} challenge{{end}}{{if existdel "autofocus"}} autofocus{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	return li.HtmlRender(htmlstr)
}
