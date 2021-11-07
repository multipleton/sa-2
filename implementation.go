package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func CalculatePostfix(input string) (string, error) {
	var stack Stack
	inputSlice := strings.Split(input, " ")
	for i := 0; i < len(inputSlice); i++ {
		currentChar := inputSlice[i]
		if currentChar == " " {
			continue
		}
		if parsed, err := strconv.ParseFloat(currentChar, 64); err == nil {
			stack.Push(parsed)
		} else {
			second, _ := stack.Pop()
			first, _ := stack.Pop()
			if calcResult, err := DoOperation(currentChar, first, second); err == false {
				stack.Push(calcResult)
			}
		}
	}
	result, _ := stack.Pop()
	return fmt.Sprintf("%f", result), nil
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
		return firstOperand / secondOperand, false
	case "^":
		return math.Pow(firstOperand, secondOperand), false
	default:
		return 0, true
	}
}
