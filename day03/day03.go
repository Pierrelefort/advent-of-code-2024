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

func SliceByMax(arr []int, value int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for pos, v := range res {
		if v > value {
			return res[:pos]
		}
	}
	return res
}

func Day03b(pathFile string) (int, error) {
	strInput, err := utils.PathToString(pathFile)
	if err != nil {
		return -1, err
	}
	regexFormulaMul := `mul\(\d{1,3},\d{1,3}\)`
	r, err := regexp.Compile(regexFormulaMul)
	if err != nil {
		return -1, err
	}
	matchesSlice := r.FindAllStringIndex(strInput, -1)

	regexFormulaDo := `do\(\)`
	rDo, _ := regexp.Compile(regexFormulaDo)
	matchesDo := rDo.FindAllStringIndex(strInput, -1)
	regexFormulaDont := `don\'t\(\)`
	rDont, _ := regexp.Compile(regexFormulaDont)
	matchesDont := rDont.FindAllStringIndex(strInput, -1)

	var doIndexList []int
	for _, slice := range matchesDo {
		doIndexList = append(doIndexList, slice[1])
	}

	var dontIndexList []int
	for _, slice := range matchesDont {
		dontIndexList = append(dontIndexList, slice[1])
	}
	sum := 0
	for _, matchSlice := range matchesSlice {
		indexMatch := matchSlice[0]
		var dontClosestIndex int
		arr := SliceByMax(dontIndexList, indexMatch)
		if len(arr) > 0 {
			dontClosestIndex = arr[len(arr)-1]
		} else {
			dontClosestIndex = 0
		}

		var doClosestIndex int
		arr = SliceByMax(doIndexList, indexMatch)
		if len(arr) > 0 {
			doClosestIndex = arr[len(arr)-1]
		} else {
			doClosestIndex = 0
		}

		validMul := true
		if doClosestIndex < dontClosestIndex {
			validMul = false
		}

		if validMul {
			currString := strInput[matchSlice[0]+4 : matchSlice[1]-1]
			values := strings.Split(currString, ",")
			if len(values) > 2 {
				fmt.Println("More than two values, unexpected match", matchSlice)
				return -1, err
			}
			intValue1, _ := strconv.Atoi(values[0])
			intValue2, _ := strconv.Atoi(values[1])
			sum += (intValue1 * intValue2)
		}

	}

	return sum, nil
}
