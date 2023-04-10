func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	retVal := 0
	for i < j {
		if height[i] > height[j] {
			if retVal < height[j]*(j-i) {
				retVal = height[j] * (j - i)
			}
			j--
		} else {
			if retVal < height[i]*(j-i) {
				retVal = height[i] * (j - i)
			}
			i++
		}
	}
	return retVal
}