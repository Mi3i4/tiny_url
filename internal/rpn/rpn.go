package rpn

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrDivByZero         = errors.New("division by zero")
	ErrNotEnoughOperands = errors.New("not enough operands")
	ErrInvalidExpr       = errors.New("invalid expression")
)

func RPNCalc(expr string) (float64, error) {
	tokens := strings.Fields(expr)
	stack := make([]float64, 0, len(tokens))

	for _, tok := range tokens {
		switch tok {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, ErrNotEnoughOperands
			}

			last := stack[len(stack)-1]
			first := stack[len(stack)-2]

			var res float64
			switch tok {
			case "+":
				res = first + last
			case "-":
				res = first - last
			case "*":
				res = first * last
			case "/":
				if last == 0 {
					return 0, ErrDivByZero
				}
				res = first / last
			}

			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		default:
			num, err := strconv.ParseFloat(tok, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		}

	}
	if len(stack) != 1 {
		return 0, ErrInvalidExpr
	}
	return stack[0], nil
}
