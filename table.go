package lispy

func Table(li *Lispy) string {
	str := `<table `

	if li.ExistDel("border") {
		str += `border="1"`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</table>`

	return str
}

func TableTd(li *Lispy) string {
	str := `<` + li.Name + ` `

	if li.Exist("colspan") {
		str += `colspan="` + li.GetDel("colspan") + `"`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</` + li.Name + `>`

	return str
}
