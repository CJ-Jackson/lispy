package lispy

func Area(li *Lispy) string {
	if !li.Exist("shape") {
		li.Set("shape", li.Content)
	}

	str := `<area`

	if li.Exist("shape") {
		str += ` shape="` + li.GetDel("shape") + `"`
	}

	if li.Exist("coords") {
		str += ` coords="` + li.GetDel("coords") + `"`
	}

	if li.Exist("href") {
		str += ` href="` + li.GetDel("href") + `"`
	}

	if li.Exist("alt") {
		str += ` alt="` + li.GetDel("alt") + `"`
	}

	str += li.GetParam()

	str += `/>`

	return str
}

func AreaNoFollow(li *Lispy) string {
	if !li.Exist("shape") {
		li.Set("shape", li.Content)
	}

	str := `<area`

	if li.Exist("shape") {
		str += ` shape="` + li.GetDel("shape") + `"`
	}

	if li.Exist("coords") {
		str += ` coords="` + li.GetDel("coords") + `"`
	}

	if li.Exist("href") {
		str += ` href="` + li.GetDel("href") + `"`
	}

	if li.Exist("alt") {
		str += ` alt="` + li.GetDel("alt") + `"`
	}

	str += ` rel="nofollow"`

	str += li.GetParam()

	str += `/>`

	return str
}
