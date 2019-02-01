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
	fname [20]rune
	lname [20]rune
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
		line := scanner.Text()
		p := person{}
		i := strings.Index(line, " ")
		copy(p.fname[:], []rune(line[:i]))
		copy(p.lname[:], []rune(line[i+1:]))
		s = append(s, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, p := range s {
		fmt.Printf("%s %s\n", string(p.fname[:]), string(p.lname[:]))
	}
}
