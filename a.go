package lispy

func A(li *Lispy) string {
	const htmlstr = `{{$li := .}}<a href="{{.GetDel "href"|html}}"{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</a>`

	if !li.Exist("href") {
		li.Set("href", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func ANoFollow(li *Lispy) string {
	const htmlstr = `{{$li := .}}<a href="{{.GetDel "href"|html}}" rel="nofollow"{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</a>`

	if !li.Exist("href") {
		li.Set("href", li.Content)
	}

	li.Delete("rel")

	return li.HtmlRender(htmlstr)
}
