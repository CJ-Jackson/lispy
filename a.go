package lispy

func A(li *Lispy) string {
	const htmlstr = `<a href="{{getdel "href"}}"{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</a>`

	if !li.Exist("href") {
		li.Set("href", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func ANoFollow(li *Lispy) string {
	const htmlstr = `<a href="{{getdel "href"}}" rel="nofollow"{{range names}} {{.}}="{{get .}}"{{end}}>{{.Content|parse}}</a>`

	if !li.Exist("href") {
		li.Set("href", li.Content)
	}

	li.Delete("rel")

	return li.HtmlRender(htmlstr)
}
