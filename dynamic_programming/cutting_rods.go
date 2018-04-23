package dynamic_programming

import "math"

func CuttingRods(price []int, length int) int {
	val := make([]int, length+1)
	val[0] = 0
	for i := 1; i < length; i++ {
		max := -1
		for j := 0; j < i; j++ {
			max = int(math.Max(float64(max), float64(price[j]+val[j-i-1])))
		}
		val[i] = max
	}
	return val[length]
}
