package lispy

func Img(li *Lispy) string {
	const htmlstr = `<img src="{{.Content}}" alt="{{getdel "alt"}}" title="{{getdel "title"}}"{{range names}} {{.|attr}}="{{get .}}"{{end}}/>`
	li.Delete("src")
	return li.HtmlRender(htmlstr)
}
