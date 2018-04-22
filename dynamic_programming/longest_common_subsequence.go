package dynamic_programming

import "math"

func LongestCommonSubsequence(a, b string) int {
	lcs := [][]int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if a[i-1] == b[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = int(math.Max(float64(lcs[i-1][j]), float64(lcs[i][j-1])))
			}
		}
	}
	return lcs[len(a)][len(b)]
}
