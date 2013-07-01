package lispy

func Video(li *Lispy) string {
	str := `<video`

	if li.Exist("src") {
		str += ` src="` + li.GetDel("src") + `"`
	}

	if li.Exist("poster") {
		str += ` poster="` + li.GetDel("poster") + `"`
	}

	if li.Exist("width") {
		str += ` width="` + li.GetDel("width") + `"`
	}

	if li.Exist("height") {
		str += ` height="` + li.GetDel("height") + `"`
	}

	if li.ExistRes("autoplay") {
		str += ` autoplay`
	}

	if li.ExistDel("controls") {
		str += ` controls`
	}

	if li.ExistRes("loop") {
		str += ` loop`
	}

	if li.ExistDel("muted") {
		str += ` muted`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</video>`

	return str
}

func Audio(li *Lispy) string {
	str := `<audio`

	if li.Exist("src") {
		str += ` src="` + li.GetDel("src") + `"`
	}

	if li.ExistRes("autoplay") {
		str += ` autoplay`
	}

	if li.ExistDel("controls") {
		str += ` controls`
	}

	if li.ExistRes("loop") {
		str += ` loop`
	}

	str += li.GetParam()

	str += `>`

	str += li.Render(li.Content)

	str += `</audio>`

	return str
}

func Source(li *Lispy) string {
	if !li.Exist("src") {
		li.Set("src", li.Content)
	}

	str := `<source`

	if li.Exist("src") {
		str += ` src="` + li.GetDel("src") + `"`
	}

	if li.Exist("media") {
		str += ` media="` + li.GetDel("media") + `"`
	}

	if li.Exist("type") {
		str += ` type="` + li.GetDel("type") + `"`
	}

	str += li.GetParam()

	str += `/>`

	return str
}

func Track(li *Lispy) string {
	if !li.Exist("src") {
		li.Set("src", li.Content)
	}

	str := `<track`

	if li.Exist("default") {
		str += ` default="` + li.GetDel("default") + `"`
	}

	if li.Exist("kind") {
		str += ` kind="` + li.GetDel("kind") + `"`
	}

	if li.Exist("label") {
		str += ` label="` + li.GetDel("label") + `"`
	}

	if li.Exist("src") {
		str += ` src="` + li.GetDel("src") + `"`
	}

	if li.Exist("srclang") {
		str += ` srclang="` + li.GetDel("srclang") + `"`
	}

	str += li.GetParam()

	str += `/>`

	return str
}
