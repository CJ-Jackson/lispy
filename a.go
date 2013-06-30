package lispy

func A(li *Lispy) string {
	str := `<a href="`

	if !li.Exist("href") {
		str += li.Content
	} else {
		str += li.GetDel("href")
	}

	str += `"`

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</a>`

	return str
}

func ANoFollow(li *Lispy) string {
	li.Delete("rel")

	str := `<a href="`

	if !li.Exist("href") {
		str += li.Content
	} else {
		str += li.GetDel("href")
	}

	str += `" rel="nofollow"`

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</a>`

	return str
}
