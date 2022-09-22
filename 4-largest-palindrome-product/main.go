package main

import (
	"fmt"
	"strconv"
)

const upperLimit = 999 * 999
const lowerLimit = 100 * 100

func reverse(word string) (result string) {
	for _, char := range word {
		result = string(char) + result
	}
	return
}

func inRange(number int) bool {
	return number >= 100 && number <= 999
}

func isMultiple(number int, multiple int) bool {
	return number % multiple == 0
}

func hasMultiplesInRange(number int) (found bool) {
	for j:= 100; j <= 999; j++ {
		if isMultiple(number, j) && inRange(number / j)  {
			found = true
			break
		}
	}
	return
}

func largestPalindrome() (palindrome int) {
	for i := upperLimit; i >= lowerLimit; i-- {
		if strconv.Itoa(i) == reverse(strconv.Itoa(i)) {
			if hasMultiplesInRange(i) {
				palindrome = i
				break
			}
		}
	}
	return
}

func main() {
	fmt.Println(largestPalindrome())
}
