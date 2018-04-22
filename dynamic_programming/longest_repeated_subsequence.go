package dynamic_programming

import "math"

func LongestRepeatedSubsequence(s string) int {
	lrs := [][]int{}
	for i := 0; i <= len(s); i++ {
		for j := 0; j <= len(s); j++ {
			lrs[i][j] = 0
		}
	}

	for i := 0; i <= len(s); i++ {
		for j := 0; j <= len(s); j++ {
			if s[i-1] == s[j-1] && i != j {
				lrs[i][j] = lrs[i][j] + 1
			} else {
				lrs[i][j] = int(math.Max(float64(lrs[i-1][j]), float64(lrs[i][j-1])))
			}
		}
	}
	return lrs[len(s)][len(s)]
}
