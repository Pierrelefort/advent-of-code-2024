package main

import (
	"advent-of-code/day01"
	"advent-of-code/day02"
	"advent-of-code/day03"
	"advent-of-code/utils"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/maps"
)

func main() {
	var err error
	var dayNumberString string

	functionMap := map[int][]func(string) (int, error){
		1: {day01.Day01a, day01.Day01b},
		2: {day02.Day02a, day02.Day02b},
		3: {day03.Day03a, day03.Day03b},
	}

	dayNumber := 3
	fmt.Printf("Enter day number (default:%d, available %v): ", dayNumber, maps.Keys(functionMap))
	_, err = fmt.Scanln(&dayNumberString)
	// Error of Scanln include '\n' so if an error occurs just go with default
	if err == nil && dayNumberString != "" {
		dayNumber, err = strconv.Atoi(dayNumberString)
		if err != nil {
			fmt.Printf("Invalid value, cannot convert %s to a number\n", dayNumberString)
		}
	}
	fmt.Printf("Launching Day%02d functions\n", dayNumber)

	functionList, ok := functionMap[dayNumber]
	if !ok {
		fmt.Printf("Missing functions for this day%02d\n", dayNumber)
		return
	}

	listVersion := []string{"a", "b"}
	for i := 0; i < len(listVersion); i++ {
		letter := listVersion[i]
		inputPath := utils.Root + fmt.Sprintf("/inputs/%02d%s.txt", dayNumber, letter)
		if _, err := os.Stat(inputPath); err != nil {
			fmt.Printf("File %s does not exist\n", inputPath)
			continue
		}

		if i >= len(functionList) {
			fmt.Printf("Missing function Day%02d%s\n", dayNumber, letter)
			continue
		}
		functionDay := functionList[i]

		result, err := functionDay(inputPath)
		fmt.Printf("Day%02d%s returned %d, %v\n", dayNumber, letter, result, err)
	}
}
