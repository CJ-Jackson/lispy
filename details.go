package lispy

func Details(li *Lispy) string {
	str := `<details `

	if li.ExistDel("open") {
		str += `open `
	}

	str += li.GetParam() + ` >`

	str += li.Render(li.Content)

	str += `</details>`

	return str
}
