package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func checks(list []int) bool {
	is_decreasing := list[0] > list[1]

	min_diff := 0
	max_diff := 4

	for i := 0; i+1 < len(list); i++ {
		j := i + 1

		iValue := list[i]
		jValue := list[j]
		diff := utils.Abs(iValue - jValue)

		if is_decreasing && iValue < jValue {
			return false
		} else if !is_decreasing && iValue > jValue {
			return false
		}

		if diff <= min_diff || diff >= max_diff {
			return false
		}
	}

	return true
}

func day02a(pathFile string) (int, error) {
	fileLines, err := utils.ReadLines(pathFile)
	if err != nil {
		return -1, err
	}

	count := 0
	for _, line := range fileLines {
		words := strings.Fields(line)
		intList, err := utils.SliceAtoi(words)
		if err != nil {
			return -1, err
		}

		if checks(intList) {
			count += 1
		}
	}

	return count, nil
}

func day02b(pathFile string) (int, error) {
	fileLines, err := utils.ReadLines(pathFile)
	if err != nil {
		return -1, err
	}

	count := 0
	for _, line := range fileLines {
		words := strings.Fields(line)
		intList, err := utils.SliceAtoi(words)
		if err != nil {
			return -1, err
		}

		if checks(intList) {
			count += 1
			continue
		}

		foundValid := false
		for i := 0; !foundValid && i < len(intList); i++ {
			localList := utils.RemoveIndex(intList, i)
			if checks(localList) {
				count += 1
				foundValid = true
			}
		}

		if !foundValid {
			localList := utils.RemoveIndex(intList, len(intList)-1)
			if checks(localList) {
				count += 1
			}
		}
	}

	return count, nil
}

func main() {
	inputPathTest := "inputs/input02a_test.txt"
	resTestA, err := day02a(inputPathTest)
	if err != nil {
		fmt.Println("Error result a:", err)
	}
	expectedResult := 2
	if resTestA != expectedResult {
		fmt.Println("Error in test result a:", resTestA)
	}

	inputPathA := "inputs/input02a.txt"
	resA, err := day02a(inputPathA)
	if err != nil {
		fmt.Println("Error result a:", err)
	}
	fmt.Println("Result a:", resA)

	inputPathTest = "inputs/input02b_test.txt"
	resTestB, err := day02b(inputPathTest)
	if err != nil {
		fmt.Println("Error result b:", err)
	}
	expectedResult = 4
	if resTestB != expectedResult {
		fmt.Println("Error in test result b:", resTestB)
	}

	inputPathB := "inputs/input02b.txt"
	resB, err := day02b(inputPathB)
	if err != nil {
		fmt.Println("Error result b:", err)
	}
	fmt.Println("Result b:", resB)

}
