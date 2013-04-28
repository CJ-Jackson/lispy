package lispy

func Details(li *Lispy) string {
	const htmlstr = `<details{{if existdel "open"}} open{{end}}>{{.Content|render}}</details>`
	return li.HtmlRender(htmlstr)
}
