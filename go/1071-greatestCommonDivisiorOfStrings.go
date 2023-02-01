func gcdOfStrings(str1 string, str2 string) string {
	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}
	for i := 1; i < len(str2); i++ {
		if equalStep(str2, i) &&
			len(str1)%(len(str2)/i) == 0 &&
			str2[:len(str2)/i] == str1[:len(str2)/i] &&
			equalStep(str1, i*len(str1)/len(str2)) {
			return str2[:len(str2)/i]
		}
	}
	return ""
}

func equalStep(str string, step int) bool {
	if len(str)%step != 0 {
		return false
	}

	l := len(str) / step
	target := str[:l]
	for i := l; i <= len(str)-l; i = i + l {
		if str[i:i+l] != target {
			return false
		}
	}
	return true
}