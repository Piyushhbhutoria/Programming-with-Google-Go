package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter file name:")
	scanner.Scan()
	fileName := scanner.Text()
	var fileHandler *os.File
	slice := make([]Person, 0, 1)

	for {
		fi, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			fmt.Println("Please Enter valid file name:")
			scanner.Scan()
			fileName = scanner.Text()
		} else {
			fileHandler = fi
			break
		}
	}
	fileScanner := bufio.NewScanner(fileHandler)
	var arr []string

	for fileScanner.Scan() {
		arr = strings.Split(fileScanner.Text(), " ")
		slice = append(slice, Person{arr[0], arr[1]})
	}

	fileHandler.Close()

	for _, person := range slice {
		fmt.Printf("%v - %v\n", person.fname, person.lname)
	}
}
