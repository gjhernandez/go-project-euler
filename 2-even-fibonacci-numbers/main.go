package main

import (
	"fmt"
)

func fiboSum(limit int) int {
	sum, previous, current := 0, 1, 2

	for current < limit {
		if current % 2 == 0 { sum += current }
		current = previous + current
		previous = current - previous
	}
	return sum
}

func main() {
	const limit = 4e6
	fmt.Println(fiboSum(limit))
}
