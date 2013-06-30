package lispy

func Button(li *Lispy) string {
	str := `<button`

	if li.Exist("value") {
		str += ` value="` + li.GetDel("value") + `"`
	}

	if li.Exist("name") {
		str += ` name="` + li.GetDel("name") + `"`
	}

	if li.Exist("type") {
		str += ` type="` + li.GetDel("type") + `"`
	}

	if li.ExistRes("disabled") {
		str += ` disabled`
	}

	if li.ExistRes("autofocus") {
		str += ` autofocus`
	}

	if li.ExistRes("formnovalidate") {
		str += ` formnovalidate`
	}

	str += li.GetParam()

	str += `/>`

	return str
}
