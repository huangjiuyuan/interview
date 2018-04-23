package dynamic_programming

func CoinChange(s []int, n int) int {
	cc := make([]int, n)
	cc[0] = 1
	for i := 0; i < len(s); i++ {
		for j := s[i]; j < n; j++ {
			cc[j] += cc[j-s[i]]
		}
	}
	return cc[n]
}
