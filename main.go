package main

import (
	"advent-of-code/day01"
	"advent-of-code/day02"
	"advent-of-code/day03"
	"advent-of-code/utils"
	"fmt"
)

func main() {
	inputPath01a := utils.Root + "/inputs/01a.txt"
	resDay01a := day01.Day01a(inputPath01a)
	fmt.Println("Result 01a:", resDay01a)

	inputPath01b := utils.Root + "/inputs/01b.txt"
	resDay01b := day01.Day01b(inputPath01b)
	fmt.Println("Result 01b:", resDay01b)

	inputPathTest := utils.Root + "/inputs/02a_test.txt"
	resTestA, err := day02.Day02a(inputPathTest)
	if err != nil {
		fmt.Println("Error result a:", err)
	}
	expectedResult := 2
	if resTestA != expectedResult {
		fmt.Println("Error in test result a:", resTestA)
	}

	inputPathA := utils.Root + "/inputs/02a.txt"
	resA, err := day02.Day02a(inputPathA)
	if err != nil {
		fmt.Println("Error result a:", err)
	}
	fmt.Println("Result a:", resA)

	inputPathTest = utils.Root + "/inputs/02b_test.txt"
	resTestB, err := day02.Day02b(inputPathTest)
	if err != nil {
		fmt.Println("Error result b:", err)
	}
	expectedResult = 4
	if resTestB != expectedResult {
		fmt.Println("Error in test result b:", resTestB)
	}

	inputPathB := utils.Root + "/inputs/02b.txt"
	resB, err := day02.Day02b(inputPathB)
	if err != nil {
		fmt.Println("Error result b:", err)
	}
	fmt.Println("Result b:", resB)

	pathFile := utils.Root + "/inputs/03a.txt"
	result, err := day03.Day03a(pathFile)
	if err != nil {
		return
	}
	fmt.Println(result)

	pathFile = utils.Root + "/inputs/03b.txt"
	result, err = day03.Day03b(pathFile)
	if err != nil {
		return
	}
	fmt.Println(result)
}
