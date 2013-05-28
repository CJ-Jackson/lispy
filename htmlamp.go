package lispy

func HtmlAmp(li *Lispy) string {
	return "&" + li.Content + ";"
}
