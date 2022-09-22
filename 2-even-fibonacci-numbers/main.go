package main

import (
	"fmt"
)

func fiboSum(limit int) (sum int) {
	previous, current := 1, 2

	for current < limit {
		if current % 2 == 0 { sum += current }
		current = previous + current
		previous = current - previous
	}
	return
}

func main() {
	const limit = 4e6
	fmt.Println(fiboSum(limit))
}
