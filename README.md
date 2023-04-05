# Calculator Project

This is a simple command line calculator project that can perform basic arithmetic operations such as addition, subtraction, multiplication, and division. It takes input from the command line and outputs the result of the operation or an error message if the input is invalid or the operation cannot be performed.

## Features

- The calculator is intentionally implemented to use only one arithmetic operation.
- The calculator is based on floating point numbers (float64)
- Any operation is space-insensitive.
- The calculator supports meaningful erros such as division by zero, float64 overflow, or syntax erros.
- Unit testing and Integration Testing are included.

## Installation

Clone the repository using the following command: git clone https://github.com/mokimokheonpark/Calculator-Project.git

## Usage

- First, navigate to the cloned repository directory.
- You can simply run the "main.go" file using the following command: go run main.go
- You can also make the executable file using the following command: go build  
  and then run the executable file using the following command: ./Calculator-Project
- You will be prompted to enter an arithmetic expression, which can include numbers, and operators (+, -, *, /).
  After entering the expression, press Enter to see the result of the operation or an error message.
- To exit the calculator, type "EXIT", "Exit", or "exit" and press Enter.

## Tests

- First, navigate to the cloned repository directory.
- Run operation_test using the following command: go test -v ./operation
- Run parser_test using the following command: go test -v ./parser

## Contributions

Contributions to the project are welcome! If you find any issues or have any suggestions for improvement, feel free to create an issue or a pull request.

## License

The project is licensed under the MIT License.
