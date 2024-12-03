package day03

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day03a(pathFile string) (int, error) {
	strInput, err := utils.PathToString(pathFile)
	if err != nil {
		return -1, err
	}
	regexFormulaMul := `mul\(\d{1,3},\d{1,3}\)`
	r, err := regexp.Compile(regexFormulaMul)
	if err != nil {
		return -1, err
	}
	matches := r.FindAllString(strInput, -1)

	sum := 0
	for _, match := range matches {
		currString := match[4 : len(match)-1]
		values := strings.Split(currString, ",")
		if len(values) > 2 {
			fmt.Println("More than two values, unexpected match", match)
			return -1, err
		}
		intValue1, _ := strconv.Atoi(values[0])
		intValue2, _ := strconv.Atoi(values[1])
		sum += (intValue1 * intValue2)
	}

	return sum, nil
}

func Day03b(pathFile string) (int, error) {
	return 0, nil
}
