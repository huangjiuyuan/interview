package dynamic_programming

import "math"

func MinCostPath(a [][]int, m, n int) int {
	mcp := [][]int{}
	mcp[0][0] = a[0][0]
	for i := 1; i < m; i++ {
		mcp[i][0] = mcp[i-1][0] + a[i][0]
	}
	for j := 1; j < n; j++ {
		mcp[0][j] = mcp[0][j-1] + a[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			min := int(math.Min(float64(mcp[i-1][j]), float64(mcp[i][j-1])))
			min = int(math.Min(float64(mcp[i-1][j-1]), float64(min)))
			mcp[i][j] = min
		}
	}
	return mcp[m-1][n-1]
}
