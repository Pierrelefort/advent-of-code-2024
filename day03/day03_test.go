package day03

import (
	"advent-of-code/utils"
	"testing"
)

func TestDay03a(t *testing.T) {
	testPath := utils.Root + "/inputs/03a_test.txt"
	expectedResult := 161

	result, err := Day03a(testPath)
	if result != expectedResult || err != nil {
		t.Fatalf(`day03a returned %d, %v, wanted %d, nil`, result, err, expectedResult)
	}
}

func TestDay03b(t *testing.T) {
	testPath := utils.Root + "/inputs/03b_test.txt"
	expectedResult := 161

	result, err := Day03a(testPath)
	if result != expectedResult || err != nil {
		t.Fatalf(`day03a returned %d, %v, wanted %d, nil`, result, err, expectedResult)
	}
}
