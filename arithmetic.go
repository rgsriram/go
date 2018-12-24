package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func performArithmetic(num1 int, num2 int, action string) int {
	switch action {

	case "add":
		return num1 + num2

	case "sub":
		return num1 - num2

	case "mul":
		return num1 * num2

	case "div":
		return num1 / num2

	default:
		return 0
	}
}

func main() {
	var operation string

	fmt.Println("Enter operation (add/sub/mul/div): ")
	fmt.Scan(&operation)

	fmt.Println("Enter numbers:")
	scanner := bufio.NewScanner(os.Stdin)

	var input [2]int
	var i int = 0

	for scanner.Scan() {

		n, err := strconv.Atoi(scanner.Text())

		if err == nil {
			input[i] = n
		}

		i++

		if i > 1 {
			break
		}

	}

	fmt.Println("Result: ", performArithmetic(input[0], input[1], operation))

}
