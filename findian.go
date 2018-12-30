package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string

	fmt.Println("Enter the String ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Not Found!")
		return
	}

	input = strings.Trim(input, " \n")

	caseInSensitiveInput := strings.ToLower(input)

	if caseInSensitiveInput == "" {
		fmt.Println("Not Found!")
		return
	}

	if caseInSensitiveInput[0] == 'i' && caseInSensitiveInput[len(caseInSensitiveInput)-1] == 'n' && strings.Contains(caseInSensitiveInput, "a") {
		fmt.Println("Found!")
		return
	}

	fmt.Println("Not Found!")
}
