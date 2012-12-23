package lispy

func Br(li *Lispy) string {
	num := li.ContentAsInt64()

	if num <= 0 {
		num = 1
	} else if num > 10 {
		num = 10
	}

	count := int64(0)
	str := ""

	for {
		if count == num {
			break
		}

		str += "<br />"
		count++
	}

	return str
}
