package greedy_algorithms

import "fmt"

func ActivitySelection(start, finish []int) {
	i := 0
	fmt.Printf("%d ", i)
	for j := 1; j < len(finish); j++ {
		if start[j] >= finish[i] {
			fmt.Printf("%d ", j)
			i = j
		}
	}
}
