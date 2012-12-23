package lispy

func Details(li *Lispy) string {
	const htmlstr = `<details{{if existdel "open"}} open{{end}}>{{.Content|parse}}</details>`
	return li.HtmlRender(htmlstr)
}
