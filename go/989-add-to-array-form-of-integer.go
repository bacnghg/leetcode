func addToArrayForm(num []int, k int) []int {
	result := []int{}
	i := len(num) - 1

	for i >= 0 || k > 0 {
		if i >= 0 {
			k = k + num[i]
		}
		result = append(result, k%10)
		i--
		k = k / 10
	}

	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}
	return result
}