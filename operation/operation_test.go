package operation_test

import (
	"testing"

	"github.com/mokimokheonpark/Calculator-Project/operation"
)

// Unit Testing of Number.Operate Function

func TestNumberOperate(t *testing.T) {
	testCases := []struct {
		value    float64
		expected float64
	}{
		{5, 5},
		{-12.345, -12.345},
		{0, 0},
	}

	for _, testCase := range testCases {
		result, _ := operation.Number{Value: testCase.value}.Operate()
		if result != testCase.expected {
			t.Errorf("Expected: %v, but got: %v", testCase.expected, result)
		}
	}
}

// Unit Testing of Addition.Operate Function

func TestAdditionOperate(t *testing.T) {
	testCases := []struct {
		leftNumber  float64
		rightNumber float64
		expected    float64
	}{
		{1, 2, 3},
		{6, -1, 5},
		{-100, 15, -85},
		{-1, -10000, -10001},
		{0, 0, 0},
		{7, 0, 7},
		{0, -2, -2},
		{-5, 5, 0},
		{1, 0.1, 1.1},
		{-0.001, 0, -0.001},
		{1.234, -5.678, -4.444},
		{-2.5, -7.5, -10},
		{5.00, 6.00000, 11},
	}

	for _, testCase := range testCases {
		result, _ := operation.Addition{LeftNumber: testCase.leftNumber, RightNumber: testCase.rightNumber}.Operate()
		if result != testCase.expected {
			t.Errorf("Expected: %v, but got: %v", testCase.expected, result)
		}
	}

	_, overflowErr := operation.Addition{LeftNumber: 1.7e+308, RightNumber: 1.2e+308}.Operate()
	if overflowErr == nil {
		t.Error("Expected overflow error but got nil")
	}

	_, underflowErr := operation.Addition{LeftNumber: -1.5e+308, RightNumber: -0.3e+308}.Operate()
	if underflowErr == nil {
		t.Error("Expected underflow error but got nil")
	}
}

// Unit Testing of Subtraction.Operate Function

func TestSubtractionOperate(t *testing.T) {
	testCases := []struct {
		leftNumber  float64
		rightNumber float64
		expected    float64
	}{
		{7, 2, 5},
		{36, -3, 39},
		{-100, 11, -111},
		{-10000, -8800, -1200},
		{0, 0, 0},
		{6, 0, 6},
		{0, -22, 22},
		{5, 5, 0},
		{2, 0.3, 1.7},
		{-0.0000009, 0, -0.0000009},
		{1.2345, -6.789, 8.0235},
		{-3.5, -6.5, 3},
		{10.00, 4.0000000000, 6},
	}

	for _, testCase := range testCases {
		result, _ := operation.Subtraction{LeftNumber: testCase.leftNumber, RightNumber: testCase.rightNumber}.Operate()
		if result != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, result)
		}
	}

	_, overflowError := operation.Subtraction{LeftNumber: 1.6e+308, RightNumber: -1.5e+308}.Operate()
	if overflowError == nil {
		t.Error("Expected overflow error, but got nil")
	}

	_, underflowError := operation.Subtraction{LeftNumber: -1.7e+308, RightNumber: 0.1e+308}.Operate()
	if underflowError == nil {
		t.Error("Expected underflow error, but got nil")
	}
}

// Unit Testing of Multiplication.Operate Function

func TestMultiplicationOperate(t *testing.T) {
	testCases := []struct {
		leftNumber  float64
		rightNumber float64
		expected    float64
	}{
		{2, 4, 8},
		{5, -3, -15},
		{-100, 12, -1200},
		{-20, -4000, 80000},
		{0, 0, 0},
		{9, 0, 0},
		{0, -123456789, 0},
		{100, 100, 10000},
		{99999999, 1, 99999999},
		{10, 0.3, 3},
		{-0.001, 400000, -400},
		{3.5, -4.5, -15.75},
		{-10.1, -0.111, 1.1211},
		{9.00, 8.00000, 72},
	}

	for _, testCase := range testCases {
		result, _ := operation.Multiplication{LeftNumber: testCase.leftNumber, RightNumber: testCase.rightNumber}.Operate()
		if result != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, result)
		}
	}

	_, overflowError1 := operation.Multiplication{LeftNumber: 1.79e+308, RightNumber: 100}.Operate()
	if overflowError1 == nil {
		t.Error("Expected overflow error, but got nil")
	}

	_, overflowError2 := operation.Multiplication{LeftNumber: -12.34, RightNumber: -1.797e+308}.Operate()
	if overflowError2 == nil {
		t.Error("Expected overflow error, but got nil")
	}

	_, underflowError1 := operation.Multiplication{LeftNumber: 1.8e+156, RightNumber: -1.8e+156}.Operate()
	if underflowError1 == nil {
		t.Error("Expected underflow error, but got nil")
	}

	_, underflowError2 := operation.Multiplication{LeftNumber: -1.0e+200, RightNumber: 1.2e+109}.Operate()
	if underflowError2 == nil {
		t.Error("Expected underflow error, but got nil")
	}
}

// Unit Testing of Division.Operate Function

func TestDivisionOperate(t *testing.T) {
	testCases := []struct {
		leftNumber  float64
		rightNumber float64
		expected    float64
	}{
		{6, 2, 3},
		{15, -3, -5},
		{-100, 8, -12.5},
		{-3, -60000, 0.00005},
		{0, 123456789, 0},
		{987654321, 1, 987654321},
		{7777777, 7777777, 1},
		{90, 0.5, 180},
		{-0.03, 200, -0.00015},
		{10.5, -3.5, -3},
		{-99.99, -6.6, 15.15},
		{125.00, 5.00000, 25},
	}

	for _, testCase := range testCases {
		result, _ := operation.Division{LeftNumber: testCase.leftNumber, RightNumber: testCase.rightNumber}.Operate()
		if result != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, result)
		}
	}

	_, divisionByZeroError1 := operation.Division{LeftNumber: 21, RightNumber: 0}.Operate()
	if divisionByZeroError1 == nil {
		t.Error("Expected division by zero error, but got nil")
	}

	_, divisionByZeroError2 := operation.Division{LeftNumber: 0, RightNumber: 0}.Operate()
	if divisionByZeroError2 == nil {
		t.Error("Expected division by zero error, but got nil")
	}
}

// Unit Testing of GetExpression Function

func TestGetExpression(t *testing.T) {
	testCases := []struct {
		operator    string
		leftNumber  float64
		rightNumber float64
		expected    operation.Expression
	}{
		{"+", 1, 2, operation.Addition{LeftNumber: 1, RightNumber: 2}},
		{"-", 9, 7, operation.Subtraction{LeftNumber: 9, RightNumber: 7}},
		{"*", 3, 5, operation.Multiplication{LeftNumber: 3, RightNumber: 5}},
		{"/", 8, 4, operation.Division{LeftNumber: 8, RightNumber: 4}},
	}

	for _, testCase := range testCases {
		result, _ := operation.GetExpression(testCase.operator, testCase.leftNumber, testCase.rightNumber)
		if result != testCase.expected {
			t.Errorf("Expected %v, but got %v", testCase.expected, result)
		}
	}

	_, invalidOperatorError1 := operation.GetExpression("^", 2, 10)
	if invalidOperatorError1 == nil {
		t.Error("Expected invalid operator error, but got nil")
	}

	_, invalidOperatorError2 := operation.GetExpression("%", 20, 6)
	if invalidOperatorError2 == nil {
		t.Error("Expected invalid operator error, but got nil")
	}
}
