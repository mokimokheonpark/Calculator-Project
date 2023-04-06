package parser_test

import (
	"reflect"
	"testing"

	"github.com/mokimokheonpark/Calculator-Project/operation"
	"github.com/mokimokheonpark/Calculator-Project/parser"
)

// Unit Testing of GetTokens Function

func TestGetTokens(t *testing.T) {

	var emptySlice []string

	testCases := []struct {
		str      string
		expected []string
	}{
		{"", emptySlice},
		{"123456", []string{"123456"}},
		{"   -7.89", []string{"-7.89"}},
		{"+     ", []string{"+"}},
		{"Ab3?", []string{"Ab3?"}},
		{"1+", []string{"1", "+"}},
		{"2 -", []string{"2", "-"}},
		{" -3  *     ", []string{"-3", "*"}},
		{"4.5/", []string{"4.5", "/"}},
		{"123 4567", []string{"123", "4567"}},
		{"+ 4567", []string{"+", "4567"}},
		{"+-", []string{"+", "-"}},
		{"ABC   abcde", []string{"ABC", "abcde"}},
		{"HelloWorld*", []string{"HelloWorld", "*"}},
		{"1 + 2", []string{"1", "+", "2"}},
		{"2--5", []string{"2", "-", "-5"}},
		{"   -3     *       -7  ", []string{"-3", "*", "-7"}},
		{"+12.48 / +2.55", []string{"+12.48", "/", "+2.55"}},
		{" ABC+123   ", []string{"ABC", "+", "123"}},
		{"812376   0.11  0", []string{"812376", "0.11", "0"}},
		{"7 +B", []string{"7", "+", "B"}},
		{"+ Z 1", []string{"+", "Z", "1"}},
		{"x - y", []string{"x", "-", "y"}},
		{"+12.1212 token *", []string{"+12.1212", "token", "*"}},
		{"1 + 2 -3", []string{"1", "+", "2", "-3"}},
		{"2** 10", []string{"2", "*", "*", "10"}},
		{"0.5 / 2.3 *", []string{"0.5", "/", "2.3", "*"}},
		{"--2--- -", []string{"-", "-", "2---", "-"}},
		{"10 + 2 + 3.7", []string{"10", "+", "2", "+", "3.7"}},
		{"   a + 10-5.7 * % A3b  / gggg  ", []string{"a", "+", "10-5.7", "*", "%", "A3b", "/", "gggg"}},
	}

	for _, testCase := range testCases {
		result := parser.GetTokens(testCase.str)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expected: %v, but got: %v", testCase.expected, result)
		}
	}
}

// Unit Testing of GetExpressionWithOneToken Function

func TestGetExpressionWithOneToken(t *testing.T) {
	testCases := []struct {
		str      string
		expected operation.Expression
		isError  bool
	}{
		{"7", operation.Number{Value: 7}, false},
		{"-3.45", operation.Number{Value: -3.45}, false},
		{"A", nil, true},
		{"+", nil, true},
		{"-123abc", nil, true},
		{"0.987d", nil, true},
	}

	for _, testCase := range testCases {
		result, err := parser.GetExpressionWithOneToken(testCase.str)
		if result != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, result)
		}
		if testCase.isError && err == nil {
			t.Errorf("Expected invalid number error, but got nil")
		}
	}
}

// Unit Testing of GetExpressionWithThreeTokens Function
func TestGetExpressionWithThreeTokens(t *testing.T) {
	testCases := []struct {
		leftNumber  string
		operator    string
		rightNumber string
		expected    operation.Expression
		isError     bool
	}{
		{"1", "+", "2", operation.Addition{LeftNumber: 1, RightNumber: 2}, false},
		{"200", "-", "13.5", operation.Subtraction{LeftNumber: 200, RightNumber: 13.5}, false},
		{"123.45", "*", "-1000000", operation.Multiplication{LeftNumber: 123.45, RightNumber: -1000000}, false},
		{"-0.01", "/", "-99.9999999", operation.Division{LeftNumber: -0.01, RightNumber: -99.9999999}, false},
		{"a", "*", "10", nil, true},
		{"3232.3232", "%", "-11.11", nil, true},
		{"-100", "+", "ABCDE", nil, true},
		{"1", "2", "3", nil, true},
		{"x", "y", "0", nil, true},
		{"A", "-", "B", nil, true},
		{"10", "hello", "/", nil, true},
		{"+", "80", "*", nil, true},
	}

	for _, testCase := range testCases {
		result, err := parser.GetExpressionWithThreeTokens(testCase.leftNumber, testCase.operator, testCase.rightNumber)
		if result != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, result)
		}
		if testCase.isError && err == nil {
			t.Errorf("Expected invalid number or operator error, but got nil")
		}
	}
}

// Unit Testing of Parser Function

func TestParser(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected operation.Expression
		isError  bool
	}{
		{[]string{"123456"}, operation.Number{Value: 123456}, false},
		{[]string{"-7.89"}, operation.Number{Value: -7.89}, false},
		{[]string{"1", "+", "2"}, operation.Addition{LeftNumber: 1, RightNumber: 2}, false},
		{[]string{"2", "-", "-5"}, operation.Subtraction{LeftNumber: 2, RightNumber: -5}, false},
		{[]string{"-3", "*", "-7"}, operation.Multiplication{LeftNumber: -3, RightNumber: -7}, false},
		{[]string{"+12.48", "/", "+2.55"}, operation.Division{LeftNumber: 12.48, RightNumber: 2.55}, false},
		{[]string{}, nil, true},
		{[]string{"+"}, nil, true},
		{[]string{"Ab3?"}, nil, true},
		{[]string{"2", "-"}, nil, true},
		{[]string{"123", "4567"}, nil, true},
		{[]string{"+", "4567"}, nil, true},
		{[]string{"ABC", "+", "123"}, nil, true},
		{[]string{"812376", "0.11", "0"}, nil, true},
		{[]string{"7", "+", "B"}, nil, true},
		{[]string{"+", "Z", "1"}, nil, true},
		{[]string{"x", "-", "y"}, nil, true},
		{[]string{"+12.1212", "token", "*"}, nil, true},
		{[]string{"*", "X", "/"}, nil, true},
		{[]string{"2", "*", "*", "10"}, nil, true},
		{[]string{"10", "+", "2", "+", "3.7"}, nil, true},
	}

	for _, testCase := range testCases {
		result, err := parser.Parser(testCase.strs)
		if result != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, result)
		}
		if err == nil && testCase.isError {
			t.Error("Expected error, but got nil")
		}
	}
}
