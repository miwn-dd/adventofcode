package main

import (
	"fmt"
	"strconv"
	"strings"
)

var inputString = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,10,23,1,23,6,27,1,6,27,31,1,13,31,35,1,13,35,39,1,39,13,43,2,43,9,47,2,6,47,51,1,51,9,55,1,55,9,59,1,59,6,63,1,9,63,67,2,67,10,71,2,71,13,75,1,10,75,79,2,10,79,83,1,83,6,87,2,87,10,91,1,91,6,95,1,95,13,99,1,99,13,103,2,103,9,107,2,107,10,111,1,5,111,115,2,115,9,119,1,5,119,123,1,123,9,127,1,127,2,131,1,5,131,0,99,2,0,14,0"

func parseCode(input string) []int {
	// parse code
	var code []int
	for _, value := range strings.Split(input, ",") {
		intValue, _ := strconv.Atoi(value)

		code = append(code, intValue)
	}

	return code
}

func processCode(input []int) []int {

	// process
	for index := 0; ; {
		opCode := input[index]

		if opCode == 99 {
			break
		}

		indexOperand1 := input[index+1]
		indexOperand2 := input[index+2]
		indexResult := input[index+3]

		switch opCode {
		case 1:
			input[indexResult] = input[indexOperand1] + input[indexOperand2]
		case 2:
			input[indexResult] = input[indexOperand1] * input[indexOperand2]
		}

		index += 4
	}

	// convert result to string
	return input
}

func computePuzzle(param1 int, param2 int) int {
	input := parseCode(inputString)
	input[1] = param1
	input[2] = param2

	return processCode(input)[0]
}

func puzzle1() {
	result := computePuzzle(12, 2)

	fmt.Printf("Result puzzle 1: %v\n", result)
}

func puzzle2() {
	for param1 := 0; param1 < 100; param1++ {
		for param2 := 0; param2 < 100; param2++ {
			result := computePuzzle(param1, param2)

			if result == 19690720 {
				fmt.Printf("Result puzzle 2: %v\n", param1*100+param2)
				break
			}
		}
	}

}

func main() {
	puzzle1()
	puzzle2()
}
