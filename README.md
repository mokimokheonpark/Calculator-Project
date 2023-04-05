# Calculator Project

The Calculator Project is a simple command-line calculator written in Go that performs basic arithmetic operations such as addition, subtraction, multiplication, and division. It takes input from the command line and outputs the result of the operation or an error message if the input is invalid or the operation cannot be performed.

## Features

The Calculator Project has the following features:

- The calculator is based on floating point numbers (float64).
- The calculator is intentionally implemented to use only one arithmetic operation.
- Any operation is space-insensitive.
- The calculator supports meaningful erros such as division by zero, float64 overflow, and syntax erros.
- The project includes unit testing and integration testing.

## Prerequisite

To run the Calculator Project, you will need:

- Go 1.16 or higher installed on your machine.

## Installation

To install the Calculator Project,

- Clone the repository using the following command: git clone https://github.com/mokimokheonpark/Calculator-Project.git

## Usage

To use the Calculator Project, follow these steps:

1. Navigate to the cloned repository directory.
2. Follow either 2-1 or 2-2.
    - 2-1. Run the "main.go" file using the following command: go run main.go
    - 2-2. Make the executable file using the following command: go build  
           and then run the executable file using the following command: ./Calculator-Project
3. You will be prompted to enter an arithmetic expression, which can include numbers, and operators (+, -, *, /).
4. After entering the expression, press Enter to see the result of the operation or an error message.
5. To exit the calculator, type "EXIT", "Exit", or "exit" and press Enter.

Here are some examples of input expressions and their expected outputs:

- 3 + 7 : 10
- -4 - -1 : -3
- 2.5 * -3 : -7.5
- -10 / -2 : 5
- 1 / 0 : Error: division by zero
- 20 % 4 : Error: invalid operator
- 6 + a : Error: invalid right number
- 100 * : Error: insufficient number of numbers or operator

## Tests

To run the tests for the Calculator Project, follow these stpes:

1. First, navigate to the cloned repository directory.
2. Run the "operation_test.go" using the following command: go test -v ./operation
3. Run the "parser_test.go" using the following command: go test -v ./parser

The "operation_test.go" covers the tests for the operation package, which performs the arithmetic operations.  
The "parser_test.go" covers the tests for the parser package, which parses the input expression.

## Contributions

Contributions to the project are welcome! If you find any issues or have any suggestions for improvement, feel free to create an issue or a pull request.

## License

The project is licensed under the MIT License.
