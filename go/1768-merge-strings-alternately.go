func mergeAlternately(word1 string, word2 string) string {
	var res []byte
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			res = append(res, word1[i])
		}

		if i < len(word2) {
			res = append(res, word2[i])
		}
	}

	return string(res)
}