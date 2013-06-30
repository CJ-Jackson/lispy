package lispy

func Header(li *Lispy) string {
	str := `<` + li.Name

	str += li.GetParam() + `>`

	str += li.Render(li.Content)

	str += `</` + li.Name + `>`

	return str
}
