package lispy

import (
	"fmt"
	"html"
)

func Common(li *Lispy) string {
	const htmlstr = `<%s{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonValue(li *Lispy) string {
	const htmlstr = `<%s{{if exist "value"}} value="{{getdel "value"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonName(li *Lispy) string {
	const htmlstr = `<%s{{if exist "name"}} name="{{getdel "name"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonCite(li *Lispy) string {
	const htmlstr = `<%s{{if exist "cite"}} cite="{{getdel "cite"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonTitle(li *Lispy) string {
	const htmlstr = `<%s{{if exist "title"}} title="{{getdel "title"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonSpan(li *Lispy) string {
	const htmlstr = `<%s{{if exist "span"}} span="{{getdel "span"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonDir(li *Lispy) string {
	const htmlstr = `<%s{{if exist "dir"}} dir="{{getdel "dir"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonCiteDateTime(li *Lispy) string {
	const htmlstr = `<%s{{if exist "cite"}} cite="{{getdel "cite"}}"{{end}}{{if exist "datetime"}} datetime="{{getdel "datetime"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonSrc(li *Lispy) string {
	const htmlstr = `<%s{{if exist "src"}} src="{{getdel "src"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}>{{.Content|render}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func Blanks(li *Lispy) string {
	return ""
}

func JavaScript(li *Lispy) string {
	const htmlstr = `<script type="text/javascript">{{.Content|js_safe}}</script>`
	li.Content = html.UnescapeString(li.Content)
	return li.HtmlRender(htmlstr)
}

func CSS(li *Lispy) string {
	const htmlstr = `<style type="text/css" scoped>{{.Content|css_safe}}</style>`
	li.Content = html.UnescapeString(li.Content)
	return li.HtmlRender(htmlstr)
}

func Canvas(li *Lispy) string {
	const htmlstr = `<canvas{{if exist "id"}} id="{{getdel "id"}}"{{end}}{{range names}} {{.|attr}}="{{get .}}"{{end}}></canvas>`

	if !li.Exist("id") {
		li.Set("id", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func ScriptFile(li *Lispy) string {
	const htmlstr = `<script src="{{.Content}}"></script>`
	return li.HtmlRender(htmlstr)
}
