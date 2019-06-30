package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	var (
		a   string
		ani Animal
		r   string
	)
	for {
		fmt.Print(">")
		fmt.Scan(&a)
		fmt.Scan(&r)
		if a == "cow" {
			ani = Animal{"grass", "walk", "moo"}
		} else if a == "bird" {
			ani = Animal{"worms", "fly", "peep"}
		} else if a == "snake" {
			ani = Animal{"mice", "slither", "hsss"}
		} else {
			fmt.Println("Incorrect input. Try again")
		}
		if r == "eat" {
			ani.Eat()
		} else if r == "move" {
			ani.Move()
		} else if r == "speak" {
			ani.Speak()
		} else {
			fmt.Println("Incorrect input. Try again")
		}
	}
}

func (animal *Animal) Eat() {
	fmt.Println("I eat", animal.food)
}

func (animal *Animal) Move() {
	fmt.Println("I", animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println("I", animal.noise)
}
