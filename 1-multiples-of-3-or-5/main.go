package main

import "fmt"

func calculateSum(limit int) (sum int) {
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
	return
}

func main() {
	const limit = 1000
	fmt.Println(calculateSum(limit))
}
