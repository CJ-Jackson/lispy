package lispy

func Video(li *Lispy) string {
	const htmlstr = `<video{{if exist "src"}} src="{{getdel "src"}}"{{end}}{{if exist "poster"}} poster="{{getdel "poster"}}"{{end}}{{if exist "width"}} width="{{getdel "width"}}"{{end}}{{if exist "height"}} height="{{getdel "height"}}"{{end}}{{if existres "autoplay"}} autoplay{{end}}{{if existdel "controls"}} controls{{end}}{{if existres "loop"}} loop{{end}}{{if existdel "muted"}} muted{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</video>`
	return li.HtmlRender(htmlstr)
}

func Audio(li *Lispy) string {
	const htmlstr = `<audio{{if exist "src"}} src="{{getdel "src"}}"{{end}}{{if existres "autoplay"}} autoplay{{end}}{{if existdel "controls"}} controls{{end}}{{if existres "loop"}} loop{{end}}{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</audio>`
	return li.HtmlRender(htmlstr)
}

func Source(li *Lispy) string {
	const htmlstr = `<source{{if exist "src"}} src="{{getdel "src"}}"{{end}}{{if exist "media"}} media="{{getdel "media"}}"{{end}}{{if exist "type"}} type="{{getdel "type"}}"{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("src") {
		li.Set("src", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func Track(li *Lispy) string {
	const htmlstr = `<track{{if exist "default"}} default="{{getdel "default"}}"{{end}}{{if exist "kind"}} kind="{{getdel "kind"}}"{{end}}{{if exist "label"}} label="{{getdel "label"}}"{{end}}{{if exist "src"}} src="{{getdel "src"}}"{{end}}{{if exist "srclang"}} srclang="{{getdel "srclang"}}"{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("src") {
		li.Set("src", li.Content)
	}

	return li.HtmlRender(htmlstr)
}
