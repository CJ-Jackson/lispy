package lispy

import (
	"fmt"
)

func Common(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{range .GetNames}} {{.|attr}}="{{$li.Get .|html}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonValue(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "value"}} value="{{.GetDel "value"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonName(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "name"}} name="{{.GetDel "name"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonCite(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "cite"}} cite="{{.GetDel "cite"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonTitle(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "title"}} title="{{.GetDel "title"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonSpan(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "span"}} span="{{.GetDel "span"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonDir(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "dir"}} dir="{{.GetDel "dir"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonCiteDateTime(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "cite"}} cite="{{.GetDel "cite"|html}}"{{end}}{{if .Exist "datetime"}} datetime="{{.GetDel "datetime"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func CommonSrc(li *Lispy) string {
	const htmlstr = `{{$li := .}}<%s{{if .Exist "src"}} src="{{.GetDel "src"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}>{{.RenderedContent}}</%s>`
	return li.HtmlRender(fmt.Sprintf(htmlstr, li.Name, li.Name))
}

func Blanks(li *Lispy) string {
	return ""
}

func JavaScript(li *Lispy) string {
	const htmlstr = `<script type="text/javascript">{{.Content|js}}</script>`
	li.Content = li.RawContent()
	return li.HtmlRender(htmlstr)
}

func CSS(li *Lispy) string {
	const htmlstr = `<style type="text/css" scoped>{{.Content|css}}</style>`
	li.Content = li.RawContent()
	return li.HtmlRender(htmlstr)
}

func Canvas(li *Lispy) string {
	const htmlstr = `{{$li := .}}<canvas{{if .Exist "id"}} id="{{.GetDel "id"|html}}"{{end}}{{range .GetNames}} {{.|attr}}="{{$li.Get .}}"{{end}}></canvas>`

	if !li.Exist("id") {
		li.Set("id", li.Content)
	}

	return li.HtmlRender(htmlstr)
}

func ScriptFile(li *Lispy) string {
	const htmlstr = `<script src="{{.Content}}"></script>`
	return li.HtmlRender(htmlstr)
}

func Escape(li *Lispy) string {
	return li.Render(li.Content)
}

func Raw(li *Lispy) string {
	return li.RawContent()
}
