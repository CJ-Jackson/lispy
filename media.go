package lispy

func Video(li *Lispy) string {
	const htmlstr = `{{$li := .}}<video{{if .Exist "src"}} src="{{.GetDel "src"|html}}"{{end}}{{if .Exist "poster"}} poster="{{.GetDel "poster"|html}}"{{end}}{{if .Exist "width"}} width="{{.GetDel "width"|html}}"{{end}}{{if .Exist "height"}} height="{{.GetDel "height"|html}}"{{end}}{{if .ExistRes "autoplay"}} autoplay{{end}}{{if .ExistDel "controls"}} controls{{end}}{{if .ExistRes "loop"}} loop{{end}}{{if .ExistDel "muted"}} muted{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</video>`
	return li.HtmlRender(htmlstr)
}

func Audio(li *Lispy) string {
	const htmlstr = `{{$li := .}}<audio{{if .Exist "src"}} src="{{.GetDel "src"|html}}"{{end}}{{if .ExistRes "autoplay"}} autoplay{{end}}{{if .ExistDel "controls"}} controls{{end}}{{if .ExistRes "loop"}} loop{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</audio>`
	return li.HtmlRender(htmlstr)
}

func Source(li *Lispy) string {
	const htmlstr = `{{$li := .}}<source{{if .Exist "src"}} src="{{.GetDel "src"|html}}"{{end}}{{if .Exist "media"}} media="{{.GetDel "media"|html}}"{{end}}{{if .Exist "type"}} type="{{.GetDel "type"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("src") {
		li.Set("src", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Track(li *Lispy) string {
	const htmlstr = `{{$li := .}}<track{{if .Exist "default"}} default="{{.GetDel "default"|html}}"{{end}}{{if .Exist "kind"}} kind="{{.GetDel "kind"|html}}"{{end}}{{if .Exist "label"}} label="{{.GetDel "label"|html}}"{{end}}{{if .Exist "src"}} src="{{.GetDel "src"|html}}"{{end}}{{if .Exist "srclang"}} srclang="{{.GetDel "srclang"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("src") {
		li.Set("src", li.Content)
	}

	return li.HtmlRender(htmlstr)
}
