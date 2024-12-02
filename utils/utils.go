package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(pathFile string) ([]string, error) {
	readFile, err := os.Open(pathFile)
	var fileLines []string

	if err != nil {
		return fileLines, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}

func RemoveIndex(arr []int, index int) []int {
	res := make([]int, 0)
	res = append(res, arr[:index]...)
	return append(res, arr[index+1:]...)
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func SliceAtoi(stringArray []string) ([]int, error) {
	intArray := make([]int, 0, len(stringArray))

	for i := 0; i < len(stringArray); i++ {
		intArrayValue, err := strconv.Atoi(stringArray[i])
		if err != nil {
			return make([]int, 0), err
		}
		intArray = append(intArray, intArrayValue)
	}
	return intArray, nil
}
