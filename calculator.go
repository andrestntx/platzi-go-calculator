package mycalculator

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(input string, operator string) int {
	values := strings.Split(input, operator)
	operator1 := parserOperator(values[0])
	operator2 := parserOperator(values[1])

	switch operator {
	case "+":
		return (operator1 + operator2)
	case "-":
		return (operator1 - operator2)
	case "*":
		return (operator1 * operator2)
	case "/":
		return (operator1 / operator2)
	default:
		return 0
	}
}

func parserOperator(input string) int {
	operator, _ := strconv.Atoi(input)

	return operator
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}
