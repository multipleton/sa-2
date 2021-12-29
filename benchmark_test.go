package lab2

import (
	"fmt"
	"testing"
)

var operators = []string{
	"+",
	"-",
	"*",
	"/",
}

type BenchmarkCase struct {
	operands   []int
	expression string
}

func BuildBenchmarkCase(count int) *BenchmarkCase {
	operands := []int{}
	expression := ""
	n := count * count * count * count
	for i := 1; i <= n; i++ {
		operands = append(operands, i)
	}
	i := 0
	j := 0
	step := 2
	for i < len(operands) {
		expression += fmt.Sprint(operands[i])
		if i == len(operands)-1 {
			expression += " " + operators[j]
		} else if i != 0 && i%(step-1) == 0 {
			expression += " " + operators[j]
			j++
			if j == len(operators) {
				j = 0
			}
			step += 2
			if i%5 == 0 {
				step = 2
			}
		}
		if i != len(operands) {
			expression += " "
		}
		i++
	}
	return &BenchmarkCase{
		operands:   operands,
		expression: expression,
	}
}

func BenchmarkCalculatePostfix(b *testing.B) {
	cases := []*BenchmarkCase{}
	for i := 2; i <= 20; i++ {
		cases = append(cases, BuildBenchmarkCase(i))
	}
	for _, entry := range cases {
		title := fmt.Sprintf("operands=%d", len(entry.operands))
		b.Run(title, func(b *testing.B) {
			CalculatePostfix(entry.expression)
		})
	}
}
