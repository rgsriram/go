package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortArray(s []int, c chan []int) {
	defer close(c)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	c <- s
}

func main() {
	var s string

	fmt.Println("Enter the size of the input array:")
	fmt.Scan(&s)

	si, err := strconv.ParseInt(s, 10, 64)

	size := int(si)

	if err != nil {
		fmt.Println("Please enter Integer. Exiting...")
		return
	}

	slice := make([]int, 0, size)

	fmt.Println("Enter the numbers to sort: ")

	var number string
	var m int

	for m = 0; m < size; m++ {
		fmt.Scan(&number)
		n, err := strconv.ParseInt(number, 10, 64)
		num := int(n)

		if err != nil {
			fmt.Println("Please enter Integer. Exiting...")
			return
		}
		slice = append(slice, num)
	}

	o := 0
	l := (size / 4)

	tmp := make([]int, 0, size)
	result := make([]int, 0, size)

	for i := 0; i < 4; i++ {
		c := make(chan []int)

		if i == 3 {
			tmp = slice[o:size]
		} else {
			tmp = slice[o:(o + l)]
		}

		go sortArray(tmp, c)
		for res := range c {
			result = append(result, append(res)...)
		}
		o += l
	}

	c := make(chan []int)
	go sortArray(result, c)
	fmt.Println(<-c)
}
