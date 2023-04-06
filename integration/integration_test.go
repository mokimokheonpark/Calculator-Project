package integration_test

import (
	"testing"

	"github.com/mokimokheonpark/Calculator-Project/parser"
)

func TestIntegration(t *testing.T) {
	testCases := []struct {
		input            string
		expected         float64
		isParserError    bool
		isOperationError bool
	}{
		{"2 + 2", 4, false, false},
		{"3 - 5", -2, false, false},
		{"2 * 3", 6, false, false},
		{"10 / 2", 5, false, false},
	}

	for _, testCase := range testCases {
		tokens := parser.GetTokens(testCase.input)
		expression, parserError := parser.Parser(tokens)
		if testCase.isParserError && parserError == nil {
			t.Error("Expected parser error, but got nil")
		}
		result, operationError := expression.Operate()
		if testCase.isOperationError && operationError == nil {
			t.Error("Expected operation error, but got nil")
		}
		if result != testCase.expected {
			t.Errorf("Expected: %v, but got: %v", testCase.expected, result)
		}
	}
}
