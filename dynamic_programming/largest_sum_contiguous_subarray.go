package dynamic_programming

func LargestSumContiguousSubarray(a []int) int {
	max, current := 0, 0
	for _, v := range a {
		current = current + v
		if current < 0 {
			current = 0
		} else if current > max {
			max = current
		}
	}
	return max
}
