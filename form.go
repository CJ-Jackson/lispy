package lispy

func Form(li *Lispy) string {
	if !li.Exist("action") {
		li.Set("action", li.Content)
	}

	str := `<form`

	if li.Exist("action") {
		str += ` action="` + li.GetDel("action") + `"`
	}

	if li.Exist("method") {
		str += ` method="` + li.GetDel("method") + `"`
	}

	if li.ExistDel("novalidate") {
		str += ` novalidate`
	}

	str += li.GetParam()

	str += `/>`

	return str
}

func Input(li *Lispy) string {
	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	str := `<input`

	if li.Exist("name") {
		str += ` name="` + li.GetDel("name") + `"`
	}

	if li.Exist("value") {
		str += ` value="` + li.GetDel("value") + `"`
	}

	if li.ExistDel("checked") {
		str += ` checked`
	}

	if li.ExistRes("disabled") {
		str += ` disabled`
	}

	if li.ExistDel("autofocus") {
		str += ` autofocus`
	}

	if li.ExistDel("formnovalidate") {
		str += ` formnovalidate`
	}

	if li.ExistDel("multiple") {
		str += ` multiple`
	}

	if li.ExistDel("readonly") {
		str += ` readonly`
	}

	if li.ExistDel("required") {
		str += ` required`
	}

	str += li.GetParam()

	str += `/>`

	return str
}

func Textarea(li *Lispy) string {
	str := `<input`

	if li.Exist("name") {
		str += ` name="` + li.GetDel("name") + `"`
	}

	if li.Exist("value") {
		str += ` value="` + li.GetDel("value") + `"`
	}

	if li.ExistRes("disabled") {
		str += ` disabled`
	}

	if li.ExistDel("autofocus") {
		str += ` autofocus`
	}

	if li.ExistDel("multiple") {
		str += ` multiple`
	}

	if li.ExistDel("readonly") {
		str += ` readonly`
	}

	if li.ExistDel("required") {
		str += ` required`
	}

	str += li.GetParam()

	str += `>`

	str += li.Content

	str += `</textarea>`

	return str
}

func Select(li *Lispy) string {
	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	str := `<select`

	if li.Exist("name") {
		str += ` name="` + li.GetDel("name") + `"`
	}

	if li.ExistRes("disabled") {
		str += ` disabled`
	}

	if li.ExistDel("autofocus") {
		str += ` autofocus`
	}

	if li.ExistDel("multiple") {
		str += ` multiple`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</select>`

	return str
}

func Option(li *Lispy) string {
	str := `<option`

	if li.Exist("value") {
		str += ` value="` + li.GetDel("value") + `"`
	}

	if li.ExistRes("disabled") {
		str += ` disabled`
	}

	if li.ExistDel("selected") {
		str += ` selected`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</option>`

	return str
}

func Output(li *Lispy) string {
	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	str := `<output`

	if li.Exist("name") {
		str += ` name="` + li.GetDel("name") + `"`
	}

	if li.Exist("for") {
		str += ` for="` + li.GetDel("for") + `"`
	}

	str += li.GetParam()

	str += `>`

	str += `</output>`

	return str
}

func Label(li *Lispy) string {
	str := `<label`

	if li.Exist("name") {
		str += ` name="` + li.GetDel("name") + `"`
	}

	if li.Exist("for") {
		str += ` for="` + li.GetDel("for") + `"`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</label>`

	return str
}

func FieldSet(li *Lispy) string {
	str := `<fieldset`

	if li.Exist("name") {
		str += ` name="` + li.GetDel("name") + `"`
	}

	if li.ExistRes("disabled") {
		str += ` disabled`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</fieldset>`

	return str
}

func Keygen(li *Lispy) string {
	if !li.Exist("name") {
		li.Set("name", li.Content)
	}

	str := `<fieldset`

	if li.Exist("name") {
		str += ` name="` + li.GetDel("name") + `"`
	}

	if li.Exist("keytype") {
		str += ` keytype="` + li.GetDel("keytype") + `"`
	}

	if li.ExistRes("disabled") {
		str += ` disabled`
	}

	if li.ExistRes("challenge") {
		str += ` challenge`
	}

	if li.ExistDel("autofocus") {
		str += ` autofocus`
	}

	str += li.GetParam()

	str += `/>`

	return str
}
