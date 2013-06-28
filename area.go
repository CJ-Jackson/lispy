package lispy

func Area(li *Lispy) string {
	const htmlstr = `{{$li := .}}<area{{if .Exist "shape"}} shape="{{.GetDel "shape"|html}}"{{end}}{{if .Exist "coords"}} coords="{{.GetDel "coords"|html}}"{{end}}{{if .Exist "href"}} href="{{.GetDel "href"|html}}"{{end}}{{if .Exist "alt"}} alt="{{.GetDel "alt"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("shape") {
		li.Set("shape", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func AreaNoFollow(li *Lispy) string {
	const htmlstr = `{{$li := .}}<area{{if .Exist "shape"}} shape="{{.GetDel "shape"|html}}"{{end}}{{if .Exist "coords"}} coords="{{.GetDel "coords"|html}}"{{end}}{{if .Exist "href"}} href="{{.GetDel "href"|html}}"{{end}}{{if .Exist "alt"}} alt="{{.GetDel "alt"|html}}"{{end}} rel="nofollow"{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}} />`

	if !li.Exist("shape") {
		li.Set("shape", li.Content)
	}

	li.Delete("rel")

	return li.HtmlRender(htmlstr)
}
