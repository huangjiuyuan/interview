package dynamic_programming

import "math"

func EditDistance(a, b string) int {
	ed := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		ed[i] = make([]int, len(b))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if i == 0 {
				ed[i][j] = j
			} else if j == 0 {
				ed[i][j] = i
			} else if a[i] == b[j] {
				ed[i][j] = ed[i-1][j-1]
			} else {
				min := int(math.Min(float64(ed[i-1][j]), float64(ed[i][j-1])))
				min = int(math.Min(float64(min), float64(ed[i-1][j-1])))
				ed[i][j] = 1 + min
			}
		}
	}
	return ed[len(a)][len(b)]
}
