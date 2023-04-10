func threeSumClosest(nums []int, target int) int {
	lenNums, result, diff := len(nums), 0, math.MaxInt32
	if lenNums > 2 {
		sort.Ints(nums)
		for i := 0; i < lenNums-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, lenNums-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					result, diff = sum, abs(sum-target)
				}
				if sum == target {
					return result
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}