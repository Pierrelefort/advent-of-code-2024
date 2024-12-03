package day01

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day01a(inputPath string) (int, error) {
	fileLines, err := utils.ReadLines(inputPath)
	if err != nil {
		fmt.Println(err)
	}

	var leftList []int
	var rightList []int
	for i := 0; i < len(fileLines); i++ {
		words := strings.Fields(fileLines[i])
		if len(words) != 2 {
			fmt.Printf("Unexpected string in input: %s\n", words)
		}

		leftListElement, _ := strconv.Atoi(words[0])
		rightListElement, _ := strconv.Atoi(words[1])
		leftList = append(leftList, leftListElement)
		rightList = append(rightList, rightListElement)

	}
	sort.Ints(leftList)
	sort.Ints(rightList)

	sum := 0
	for i := 0; i < len(fileLines); i++ {
		currDiff := leftList[i] - rightList[i]
		sum += utils.Abs(currDiff)
	}

	return sum, nil
}

func Day01b(inputPath string) (int, error) {
	fileLines, err := utils.ReadLines(inputPath)
	if err != nil {
		fmt.Println(err)
	}

	var leftList []int
	rightListHistogram := make(map[int]int)

	for i := 0; i < len(fileLines); i++ {
		words := strings.Fields(fileLines[i])
		if len(words) != 2 {
			fmt.Printf("Unexpected string in input: %s\n", words)
		}

		leftListElement, _ := strconv.Atoi(words[0])
		rightListElement, _ := strconv.Atoi(words[1])
		leftList = append(leftList, leftListElement)

		rightListHistogram[rightListElement] += 1
	}
	sort.Ints(leftList)

	sum := 0
	for i := 0; i < len(fileLines); i++ {
		currValue := leftList[i] * rightListHistogram[leftList[i]]
		sum += currValue
	}

	return sum, nil
}
