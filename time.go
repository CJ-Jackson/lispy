package lispy

func Time(li *Lispy) string {
	str := `<` + li.Name + ` `

	if li.Exist("datetime") {
		str += `datetime="` + li.GetDel("datetime") + `" `
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</` + li.Name + `>`

	return str
}
