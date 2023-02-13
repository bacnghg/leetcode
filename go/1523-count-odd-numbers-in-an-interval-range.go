func countOdds(low int, high int) int {
	var count = 0
	for i := low; i <= high; i++ {
		if i%2 == 1 {
			count++
		}
	}
	return count
}