package calculation

// Number: struct and method

type Number struct {
	Value float64
}

func (n Number) Calculate() float64 {
	return n.Value
}

// Addition: struct and method

type Addition struct {
	LeftNumber  float64
	RightNumber float64
}

func (a Addition) Calculate() float64 {
	return a.LeftNumber + a.RightNumber
}

// Subtraction: struct and method

type Subtraction struct {
	LeftNumber  float64
	RightNumber float64
}

func (s Subtraction) Calculate() float64 {
	return s.LeftNumber - s.RightNumber
}

// Multiplication: struct and method

type Multiplication struct {
	LeftNumber  float64
	RightNumber float64
}

func (m Multiplication) Calculate() float64 {
	return m.LeftNumber * m.RightNumber
}

// Division: struct and method

type Division struct {
	LeftNumber  float64
	RightNumber float64
}

func (d Division) Calculate() float64 {
	return d.LeftNumber / d.RightNumber
}

// Expression: interface of methods

type Expression interface {
	Calculate() float64
}
