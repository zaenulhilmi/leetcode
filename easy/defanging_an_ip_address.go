package easy

func DefangIPAddress(input string) interface{} {
	var result string
	for _, v := range input {
		if v == '.' {
			result += "[.]"
		} else {
			result += string(v)
		}
	}
	return result
}

