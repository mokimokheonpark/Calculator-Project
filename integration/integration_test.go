package integration_test

import (
	"testing"

	"github.com/mokimokheonpark/Calculator-Project/parser"
)

// Integration Testing

func TestIntegration(t *testing.T) {
	testCases := []struct {
		input            string
		expected         float64
		isParserError    bool
		isOperationError bool
	}{
		// one number
		{"5", 5, false, false},
		{"-12.345", -12.345, false, false},
		{"0", 0, false, false},

		// addition
		{"1 + 2", 3, false, false},
		{"6 + -1", 5, false, false},
		{"-100 + 15", -85, false, false},
		{"-1 + -10000", -10001, false, false},
		{"0 + 0", 0, false, false},
		{"7 + 0", 7, false, false},
		{"0 + -2", -2, false, false},
		{"-5 + 5", 0, false, false},
		{"1 + 0.1", 1.1, false, false},
		{"-0.001 + 0", -0.001, false, false},
		{"1.234 + -5.678", -4.444, false, false},
		{"-2.5 + -7.5", -10, false, false},
		{"5.00 + 6.00000", 11, false, false},

		// subtraction
		{"7 - 2", 5, false, false},
		{"36 - -3", 39, false, false},
		{"-100 - 11", -111, false, false},
		{"-10000 - -8800", -1200, false, false},
		{"0 - 0", 0, false, false},
		{"6 - 0", 6, false, false},
		{"0 - -22", 22, false, false},
		{"5 - 5", 0, false, false},
		{"2 - 0.3", 1.7, false, false},
		{"-0.0000009 - 0", -0.0000009, false, false},
		{"1.2345 - -6.789", 8.0235, false, false},
		{"-3.5 - -6.5", 3, false, false},
		{"10.00 - 4.0000000000", 6, false, false},

		// multiplication
		{"2 * 4", 8, false, false},
		{"5 * -3", -15, false, false},
		{"-100 * 12", -1200, false, false},
		{"-20 * -4000", 80000, false, false},
		{"0 * 0", 0, false, false},
		{"9 * 0", 0, false, false},
		{"0 * -123456789", 0, false, false},
		{"100 * 100", 10000, false, false},
		{"99999999 * 1", 99999999, false, false},
		{"10 * 0.3", 3, false, false},
		{"-0.001 * 400000", -400, false, false},
		{"3.5 * -4.5", -15.75, false, false},
		{"-10.1 * -0.111", 1.1211, false, false},
		{"9.00 * 8.00000", 72, false, false},

		// division
		{"6 / 2", 3, false, false},
		{"15 / -3", -5, false, false},
		{"-100 / 8", -12.5, false, false},
		{"-3 / -60000", 0.00005, false, false},
		{"0 / 123456789", 0, false, false},
		{"987654321 / 1", 987654321, false, false},
		{"7777777 / 7777777", 1, false, false},
		{"90 / 0.5", 180, false, false},
		{"-0.03 / 200", -0.00015, false, false},
		{"10.5 / -3.5", -3, false, false},
		{"-99.99 / -6.6", 15.15, false, false},
		{"125.00 / 5.00000", 25, false, false},

		// division by zero error
		{"9 / 0", 0, false, true},
		{"21 / 0", 0, false, true},

		// empty input error
		{"", 0, true, true},

		// invalid number error
		{"Ab3?", 0, true, true},

		// insufficient number of numbers or operator error
		{"2 -", 0, true, true},
		{"123 4567", 0, true, true},
		{"+ 4567", 0, true, true},

		// too many numbers or operators error
		{"2 * * 10", 0, true, true},
		{"10 + 2 + 3.7", 0, true, true},

		// invalid left number error
		{"+", 0, true, true},
		{"ABC + 123", 0, true, true},
		{"+ Z 1", 0, true, true},
		{"x - y", 0, true, true},
		{"* X /", 0, true, true},

		// invalid right number error
		{"7 + B", 0, true, true},
		{"+12.1212 token *", 0, true, true},

		// invalid operator error
		{"2 ^ 10", 0, true, true},
		{"20 % 6", 0, true, true},
		{"812376 0.11 0", 0, true, true},
	}

	for _, testCase := range testCases {
		tokens := parser.GetTokens(testCase.input)
		expression, parserError := parser.Parser(tokens)
		if testCase.isParserError {
			if parserError == nil {
				t.Error("Expected parser error, but got nil")
			}
			continue
		}

		result, operationError := expression.Operate()
		if testCase.isOperationError {
			if operationError == nil {
				t.Error("Expected operation error, but got nil")
			}
			continue
		}

		if result != testCase.expected {
			t.Errorf("Expected: %v, but got: %v", testCase.expected, result)
		}
	}
}
