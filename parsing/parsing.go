package parsing

import (
	"errors"
	"strconv"
	"strings"
	"github.com/mokimokheonpark/calculator_project2/calculation"
)

// function which handles only four valid operations

func Operation(operator string, leftNum float64, rightNum float64) (Expression, error) {
	switch operator {
	case "+":
		return Addition{LeftNumber: leftNum, RightNumber: rightNum}, nil
	case "-":
		return Subtraction{LeftNumber: leftNum, RightNumber: rightNum}, nil
	case "*":
		return Multiplication{LeftNumber: leftNum, RightNumber: rightNum}, nil
	case "/":
		return Division{LeftNumber: leftNum, RightNumber: rightNum}, nil
	default:
		return nil, errors.New("invalid operator")
	}
}

// funcion which checks input validation using parse tree

func ParseTree(str string) (Expression, error) {

	chars := []rune(str)
	var token string
	var tokens []string
	var operators = "+-*/"

	// iterate over characters in argument str
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])

		// when char is one of valid operators
		if strings.Contains(operators, char) {
			if len(token) > 0 && len(tokens) == 0 {
				tokens = append(tokens, token)
				token = ""
				tokens = append(tokens, char)
			} else if len(tokens) == 1 {
				tokens = append(tokens, char)
			} else {
				token += char
			}

			// when char is space
		} else if char == " " {
			if len(token) > 0 {
				tokens = append(tokens, token)
				token = ""
			}

			// when char is any other character
		} else {
			token += char
		}
	}

	// append the last token to tokens if it is non-empty
	if len(token) > 0 {
		tokens = append(tokens, token)
	}

	switch len(tokens) {

	// the number of token is 0
	case 0:
		return nil, errors.New("input is empty")

	// the number of token is 1
	case 1:
		number, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			return nil, errors.New("invalid number")
		}

		return Number{Value: number}, nil

	// the number of token is 2
	case 2:
		return nil, errors.New("insufficient number of numbers or operator")

	// the number of token is 3
	case 3:
		// get the first token
		leftNumber, err := strconv.ParseFloat(tokens[0], 64)
		// check if the first token is valid
		if err != nil {
			return nil, errors.New("invalid left number")
		}

		// get the third token
		rightNumber, err := strconv.ParseFloat(tokens[2], 64)
		// check if the third token is valid
		if err != nil {
			return nil, errors.New("invalid right number")
		}

		// get the corressponding operation
		expression, err := Operation(tokens[1], leftNumber, rightNumber)
		// check if the second token is valid
		if err != nil {
			return nil, err
		}

		// handle the case of division by zero
		if tokens[1] == "/" && rightNumber == 0 {
			return nil, errors.New("division by zero")
		}
		
		return expression, nil

	// the number of token is more than 3
	default:
		return nil, errors.New("too many numbers or operators")
	}
}
