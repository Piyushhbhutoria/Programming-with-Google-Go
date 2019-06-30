package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("food: grass")
}

func (cow Cow) Move() {
	fmt.Println("locomotion: walk")
}

func (cow Cow) Speak() {
	fmt.Println("noise: moo")
}

type Bird struct{}

func (bird Bird) Eat() {
	fmt.Println("food: worms")
}

func (bird Bird) Move() {
	fmt.Println("locomotion: fly")
}

func (bird Bird) Speak() {
	fmt.Println("noise: peep")
}

type Snake struct{}

func (snake Snake) Eat() {
	fmt.Println("food: mice")
}

func (snake Snake) Move() {
	fmt.Println("locomotion: slither")
}

func (snake Snake) Speak() {
	fmt.Println("noise: hsss")
}

func main() {
	var req string
	var s1 string
	var s2 string
	m := make(map[string]Animal)
	cow := Cow{}
	bird := Bird{}
	snake := Snake{}

	for {
		fmt.Print(">")
		fmt.Scan(&req)
		fmt.Scan(&s1)
		fmt.Scan(&s2)

		if (req == "newanimal") {
			if s2 == "cow" {
				m[s1] = cow
			} else if s2 == "bird" {
				m[s1] = bird
			} else if s2 == "snake" {
				m[s1] = snake
			} else {
				fmt.Println("Incorrect input. Try again")
				continue
			}
			fmt.Println("Created it!")
		} else if (req == "query") {
			ani, ok := m[s1]
			if !ok {
				fmt.Println("Incorrect input. Try again")
			} else {
				if s2 == "eat" {
					ani.Eat()
				} else if s2 == "move" {
					ani.Move()
				} else if s2 == "speak" {
					ani.Speak()
				} else {
					fmt.Println("Incorrect input. Try again")
				}
			}
		} else {
			fmt.Println("Incorrect input. Try again")
		}
	}
}
