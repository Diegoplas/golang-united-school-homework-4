package string_sum

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf
const (
	ASCIIPlus  = 43
	ASCIIMinus = 45
	ASCIIZero  = 48
	ASCIINine  = 57
)

func StringSum(input string) (output string, err error) {
	firstNumber := ""
	secondNumber := ""
	input = strings.ReplaceAll(input, " ", "")
	if input == "" {
		return "", errorEmptyInput
	}

	err = validateOnlyValidCharacters(input)
	if err != nil {
		return "", err
	}
	err = validateOnlyTwoNums(input)
	if err != nil {
		return "", err
	}

	lastCharInput := len(input) - 1
	for idx := lastCharInput; idx >= 0; idx-- {
		if input[idx] == ASCIIMinus || input[idx] == ASCIIPlus {
			secondNumber = input[idx:]
			firstNumber = input[:idx]
			break
		}
	}

	total, err := sumStrings(firstNumber, secondNumber)
	if err != nil {
		return "", err
	}
	return total, nil
}

func validateOnlyTwoNums(input string) error {
	signCounter := 0
	if input[0] != ASCIIPlus && input[0] != ASCIIMinus {
		signCounter += 1
	}
	for idx := range input {
		if input[idx] == ASCIIPlus || input[idx] == ASCIIMinus {
			signCounter += 1
		}
		if signCounter > 2 {
			return errorNotTwoOperands
		}
	}
	if signCounter == 1 {
		return errorNotTwoOperands
	}

	return nil
}

func validateCharacter(char byte) bool {
	if (char >= ASCIIZero && char <= ASCIINine) || char == ASCIIPlus || char == ASCIIMinus {
		return true
	}
	return false
}

func validateOnlyValidCharacters(number string) error {
	for idx := range number {
		if !validateCharacter(number[idx]) {
			return errorNotTwoOperands
		}
	}
	return nil
}

func sumStrings(firstNumber, secondNumber string) (string, error) {
	firstNumInt, err := strconv.Atoi(firstNumber)
	if err != nil {
		return "", errorNotTwoOperands
	}
	secondNumInt, err := strconv.Atoi(secondNumber)
	if err != nil {
		return "", errorNotTwoOperands
	}
	total := strconv.Itoa(firstNumInt + secondNumInt)
	return total, nil
}
