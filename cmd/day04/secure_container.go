package main

import (
	"fmt"
)

const minPassword = 367479
const maxPassword = 893698

func passwordValid(password int) bool {
	hasDoubleDigit := false

	lastDigit := password % 10

	for password /= 10; password > 0; password /= 10 {
		digit := password % 10

		if digit > lastDigit {
			return false
		}

		if digit == lastDigit {
			hasDoubleDigit = true
		}

		lastDigit = digit
	}

	return hasDoubleDigit
}

func passwordValid2(password int) bool {
	digitCount := make(map[int]int)

	lastDigit := password % 10
	digitCount[lastDigit]++

	for password /= 10; password > 0; password /= 10 {
		digit := password % 10

		if digit > lastDigit {
			return false
		}

		digitCount[digit]++
		lastDigit = digit
	}

	for _, count := range digitCount {
		if count == 2 {
			return true
		}
	}

	return false
}

func puzzle1() int {
	validPasswords := 0

	for password := minPassword; password <= maxPassword; password++ {
		if passwordValid(password) {
			validPasswords++
		}
	}

	return validPasswords
}

func puzzle2() int {
	validPasswords := 0

	for password := minPassword; password <= maxPassword; password++ {
		if passwordValid2(password) {
			validPasswords++
		}
	}

	return validPasswords
}

func main() {
	fmt.Printf("Solution for Puzzle 1: %v\n", puzzle1())
	fmt.Printf("Solution for Puzzle 2: %v\n", puzzle2())
}
