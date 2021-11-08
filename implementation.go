package lab2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func CalculatePostfix(input string) (float64, error) {
	var stack Stack
	if len(input) == 0 {
		return 0, errors.New("invalid length")
	}
	inputSlice := strings.Split(input, " ")
	for i := 0; i < len(inputSlice); i++ {
		currentChar := inputSlice[i]
		if parsed, err := strconv.ParseFloat(currentChar, 64); err == nil {
			stack.Push(parsed)
		} else {
			second, err := stack.Pop()
			if err {
				return 0, errors.New("invalid characters")
			}
			first, err := stack.Pop()
			if err {
				return 0, errors.New("invalid characters")
			}
			if calcResult, err := DoOperation(currentChar, first, second); !err {
				stack.Push(calcResult)
			} else {
				return 0, errors.New("error with calculations")
			}
		}
	}
	result, _ := stack.Pop()
	return result, nil
}

func DoOperation(operation string, firstOperand, secondOperand float64) (float64, bool) {
	switch operation {
	case "+":
		return firstOperand + secondOperand, false
	case "-":
		return firstOperand - secondOperand, false
	case "*":
		return firstOperand * secondOperand, false
	case "/":
		if secondOperand != 0. {
			return firstOperand / secondOperand, false
		}
		return 0, true
	case "^":
		return math.Pow(firstOperand, secondOperand), false
	default:
		return 0, true
	}
}
