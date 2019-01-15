package main

import (
	"fmt"
	"strconv"
)

func Swap(ints []int64, i int) {
	ints[i], ints[i + 1] = ints[i + 1], ints[i]
}

func BubbleSort(s []int64) {
	n := len(s)

	for k := 1; k < n; k++ {
		flag := false

		for i := 0; i < n-k; i++ {
			j := i + 1
			if s[i] > s[j] {
				Swap(s, i)
				flag = true
			}
		}

		if flag == false {
			break
		}
	}
}

func main() {
	size := 10
	slice := make([]int64, 0, size)

	fmt.Println("Enter the 10 numbers to sort: ")

	var number string

	for m := 0; m < size; m++ {
		fmt.Scan(&number)
		num, err := strconv.ParseInt(number, 10, 64)

		if err != nil {
			fmt.Println("Please enter Integer. Exiting...")
			return
		}

		slice = append(slice, num)
	}

	BubbleSort(slice)
	fmt.Println(slice)
}
