package lispy

func Details(li *Lispy) string {
	const htmlstr = `<details{{if .ExistDel "open"}} open{{end}}>{{.RenderedContent}}</details>`
	return li.HtmlRender(htmlstr)
}
