package main

import (
	"fmt"
)

func largestPrimeFactor(number int) int {
	div := number
	var factors []int
	for div != 1 {
		initialDiv := div
		for i := 2; i <= div / 2 ; i++ {
			if div % i == 0 {
				factors  = append(factors, i, div/ i)
				div = div / i
				break
			}
		}
		if initialDiv == div {
			break
		}
	}

	if div == number {
		return number
	}

	return factors[len(factors)-1]
}

func main() {
	const number = 600851475143
	fmt.Println(largestPrimeFactor(number))
}
