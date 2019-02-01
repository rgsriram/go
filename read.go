package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {

	s := make([]person, 0, 0)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the file name [full path]: ")
	name, err := reader.ReadString('\n')
	name = strings.Trim(name, " \n") 

	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		s = append(s, person{fname: line[0], lname: line[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, p := range s {
		fmt.Printf("Firstname: %s, Lastname: %s\n", p.fname, p.lname)
	}
}
