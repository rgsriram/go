package main

import "fmt"

func main() {
	var floatNumber float64

	fmt.Println("Enter number to convert into int: ")
	fmt.Scan(&floatNumber)

	intNum := int(floatNumber)

	fmt.Println("Truncated number: ", intNum)
}
