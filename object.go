package lispy

func Object(li *Lispy) string {
	str := `<object `

	if li.Exist("width") {
		str += `width="` + li.GetDel("width") + `" `
	}

	if li.Exist("height") {
		str += `height="` + li.GetDel("height") + `" `
	}

	if li.Exist("data") {
		str += `data="` + li.GetDel("data") + `" `
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</object>`

	return str
}

func Param(li *Lispy) string {
	str := `<param `

	if li.Exist("name") {
		str += `name="` + li.GetDel("name") + `" `
	}

	if li.Exist("value") {
		str += `value="` + li.GetDel("value") + `" `
	}

	str += li.GetParam()

	str += `/>`

	return str
}

func Embed(li *Lispy) string {
	str := `<embed `

	if li.Exist("src") {
		str += `src="` + li.GetDel("src") + `" `
	}

	if li.Exist("type") {
		str += `type="` + li.GetDel("type") + `" `
	}

	if li.Exist("width") {
		str += `width="` + li.GetDel("width") + `" `
	}

	if li.Exist("height") {
		str += `height="` + li.GetDel("height") + `" `
	}

	str += li.GetParam()

	str += `/>`

	return str
}
