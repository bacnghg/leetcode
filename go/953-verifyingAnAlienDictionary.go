func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}

	return sort.SliceIsSorted(words, func(i, j int) bool {
		x, y := words[i], words[j]
		minLength := len(x)
		if len(y) < len(x) {
			minLength = len(y)
		}

		for z := 0; z < minLength; z++ {
			if x[z] != y[z] {
				return orderMap[x[z]] < orderMap[y[z]]
			}
		}
		return len(x) <= len(y)
	})
}