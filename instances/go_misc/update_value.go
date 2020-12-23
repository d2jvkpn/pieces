package rover

func UpdateValue(x *string, mp map[string]string) {
	if y, ok := mp[*x]; ok {
		*x = y
	}
}
