package lispy

func Img(li *Lispy) string {
	const htmlstr = `{{$li := .}}<img src="{{.Content|html}}" alt="{{.GetDel "alt"|html}}" title="{{.GetDel "title"|html}}"{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}/>`
	li.Delete("src")
	return li.HtmlRender(htmlstr)
}
