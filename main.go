package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/mokimokheonpark/calculator_project2/parsing"
)

// main function

func main() {

	// main loop
	for true {

		// input from command line
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// conditions of loop termination
		if input == "exit" || input == "Exit" || input == "EXIT" {
			break
		}

		// make a parse tree using input
		equation, err := ParseTree(input)

		// if an error exists in the parse tree, the corresponding error will be printed
		if err != nil {
			fmt.Println("Error:", err)
			// otherwise, the calculated value will be printed
		} else {
			fmt.Println(equation.Calculate())
		}

	}
}
