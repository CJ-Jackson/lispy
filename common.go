package lispy

import (
	"html"
)

func commonCodeReuse(li *Lispy, names ...string) string {
	str := `<` + li.Name

	for _, name := range names {
		if li.Exist(name) {
			str += ` ` + name + `="` + li.GetDel(name) + `"`
		}
	}

	str += li.GetParam() + `>`

	str += li.Render(li.Content)

	str += `</` + li.Name + `>`

	return str
}

func Common(li *Lispy) string {
	return commonCodeReuse(li)
}

func CommonValue(li *Lispy) string {
	return commonCodeReuse(li, "value")
}

func CommonName(li *Lispy) string {
	return commonCodeReuse(li, "name")
}

func CommonCite(li *Lispy) string {
	return commonCodeReuse(li, "cite")
}

func CommonTitle(li *Lispy) string {
	return commonCodeReuse(li, "title")
}

func CommonSpan(li *Lispy) string {
	return commonCodeReuse(li, "span")
}

func CommonDir(li *Lispy) string {
	return commonCodeReuse(li, "dir")
}

func CommonCiteDateTime(li *Lispy) string {
	return commonCodeReuse(li, "cite", "datetime")
}

func CommonSrc(li *Lispy) string {
	return commonCodeReuse(li, "src")
}

func Blanks(li *Lispy) string {
	return ""
}

func JavaScript(li *Lispy) string {
	str := `<script type="text/javascript">`

	str += li.RawContent()

	str += `</script>`

	return str
}

func CSS(li *Lispy) string {
	str := `<style type="text/css" scoped>`

	str += li.RawContent()

	str += `</style>`

	return str
}

func Canvas(li *Lispy) string {
	str := `<canvas`

	if li.Exist("id") {
		str += ` id="` + li.GetDel("id") + `"`
	}

	str += li.GetParam() + `>`

	str += `</canvas>`

	return str
}

func ScriptFile(li *Lispy) string {
	str := `<script src="`

	str += li.Content

	str += `"></script>`

	return str
}

func Escape(li *Lispy) string {
	return li.Render(li.Content)
}

func Raw(li *Lispy) string {
	return li.RawContent()
}

func EscapedHtml(li *Lispy) string {
	return html.EscapeString(li.Render(li.Content))
}
