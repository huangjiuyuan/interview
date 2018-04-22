package dynamic_programming

func LongestIncreasingSubsequence(a []int) int {
	lis := make([]int, len(a))
	for i := range lis {
		lis[i] = 1
	}

	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}

	max := 1
	for _, v := range lis {
		if v > max {
			max = v
		}
	}
	return max
}
