package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := make([]int, 0, 3)

	for {
		var numberString string

		fmt.Println("Enter the number to add: ")
		fmt.Scan(&numberString)

		numberString = strings.ToLower(numberString)

		if numberString == "x" {
			os.Exit(0)
		}

		num, err := strconv.Atoi(numberString)

		if err != nil {
			fmt.Println("You have entered wrong input")
			os.Exit(1)
		}

		s = append(s, num)

		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})

		fmt.Println(s)
	}

}
