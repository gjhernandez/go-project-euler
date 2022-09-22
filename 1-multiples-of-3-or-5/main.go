package main

import "fmt"

func calculateSum(limit int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		switch {
		case i % 15 == 0:
			sum += i
		case i % 3 == 0:
			sum += i
		case i % 5 == 0:
			sum += i
		}
	}
	return sum
}

func main() {
	const limit = 1000
	fmt.Println(calculateSum(limit))
}
