package calculation_test

import (
    "testing"
    calculation "calculator/calculation"
    "github.com/stretchr/testify/assert"
)

func TestCalculateFails(t *testing.T) {
    division := calculation.Division { LeftNumber: 10, RightNumber: 0 } 
    value, err := division.Calculate()

    assert.Equal(t, value, nil)
    assert.EqualError(t, err, "invalid operation: division by zero")
}

