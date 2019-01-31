package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, locomotion, noise string
}

type Bird struct {
	name, food, locomotion, noise string
}

type Snake struct {
	name, food, locomotion, noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func printActivity(an Animal, action string) {

	switch action {

	case "eat":
		an.Eat()

	case "move":
		an.Move()

	case "speak":
		an.Speak()

	default:
		fmt.Println("Enter valid action")
	}

}

func PrintAction(a Animal, action string) {
	switch a.(type) {

	case Cow:
		printActivity(a, action)

	case Snake:
		printActivity(a, action)

	case Bird:
		printActivity(a, action)

	}
}

func main() {

	a := make(map[string]Animal)

	for {
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		line := scanner.Text()

		s := strings.Split(line, " ")

		if len(s) < 3 {
			fmt.Println("Invalid input.")
			return
		}

		input1, input2, input3 := s[0], s[1], s[2]

		switch input1 {

		case "newanimal":

			var animalVal Animal

			switch input3 {
			case "cow":
				animalVal = Cow{input2, "grass", "walk", "moo"}

			case "bird":
				animalVal = Bird{input2, "worms", "fly", "peep"}

			case "snake":
				animalVal = Snake{input2, "mice", "slither", "hss"}

			default:
				fmt.Println("Please give valid animal category")
				return
			}

			a[input2] = animalVal

			fmt.Println("Created it!")

		case "query":

			animal, ok := a[input2]

			if ok == false {
				fmt.Println("Animal Not found")
				return
			}

			var animalVal Animal

			animalVal = animal

			switch input3 {

			case "eat":
				animalVal.Eat()

			case "move":
				animalVal.Move()

			case "speak":
				animalVal.Speak()

			default:
				fmt.Println("Enter valid action")
			}

		}
	}
}
