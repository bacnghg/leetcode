func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxVal := candies[0]
	for _, each := range candies {
		maxVal = max(maxVal, each)
	}
	result := make([]bool, len(candies))
	for idx, each := range candies {
		if each+extraCandies >= maxVal {
			result[idx] = true
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}