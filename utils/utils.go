package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "..")
)

func ReadLines(pathFile string) ([]string, error) {
	readFile, err := os.Open(pathFile)
	if err != nil {
		return make([]string, 0), err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}

func PathToString(pathFile string) (string, error) {
	b, err := os.ReadFile(pathFile)
	if err != nil {
		return "", err
	}
	return string(b), nil
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
