package day02

import (
	"advent-of-code/utils"
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

func Day02a(pathFile string) (int, error) {
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

func Day02b(pathFile string) (int, error) {
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
