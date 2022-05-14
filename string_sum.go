package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
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

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return "", fmt.Errorf("StringSum@len(input) %w", errorEmptyInput)
	}

	nums := make([]int, 0)
	for one, next := "", input; len(next) > 0; {
		one, next = cutOperand(next)

		one = strings.ReplaceAll(one, " ", "")
		one = strings.ReplaceAll(one, "+", "")

		n, cErr := strconv.Atoi(one)
		if cErr != nil {
			return "", fmt.Errorf("StringSum@Atoi %w", cErr)
		}

		nums = append(nums, n)
	}

	if len(nums) != 2 {
		return "", fmt.Errorf("StringSum@len(nums) %w", errorNotTwoOperands)
	}

	return fmt.Sprint(nums[0] + nums[1]), nil
}

func cutOperand(input string) (one string, next string) {
	const opers = "+-"
	var found bool
	for i, v := range input {
		if strings.ContainsRune(opers, v) {
			if !found {
				continue
			}
			return input[:i], input[i:]
		}

		if !unicode.IsSpace(v) {
			found = true
		}
	}

	return input, ""
}
