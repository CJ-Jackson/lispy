package lispy

func Area(li *Lispy) string {
	const htmlstr = `<area{{if exist "shape"}} shape="{{getdel "shape"}}"{{end}}{{if exist "coords"}} coords="{{getdel "coords"}}"{{end}}{{if exist "href"}} href="{{getdel "href"}}"{{end}}{{if exist "alt"}} alt="{{getdel "alt"}}"{{end}}{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("shape") {
		li.Set("shape", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func AreaNoFollow(li *Lispy) string {
	const htmlstr = `<area{{if exist "shape"}} shape="{{getdel "shape"}}"{{end}}{{if exist "coords"}} coords="{{getdel "coords"}}"{{end}}{{if exist "href"}} href="{{getdel "href"}}"{{end}}{{if exist "alt"}} alt="{{getdel "alt"}}"{{end}} rel="nofollow"{{range names}} {{.}}="{{get .}}"{{end}} />`

	if !li.Exist("shape") {
		li.Set("shape", li.Content)
	}

	li.Delete("rel")

	return li.HtmlRender(htmlstr)
}
