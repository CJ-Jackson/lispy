package lispy

func Img(li *Lispy) string {
	li.Delete("src")
	str := `<img src="`

	str += li.Content

	str += `" alt="`

	str += li.GetDel("alt")

	str += `" title="`

	str += li.GetDel("title")

	str += `"`

	str += li.GetParam()

	str += `/>`

	return str
}
