package main

import "fmt"

func main() {
	var first string
	fmt.Scanln(&first)
	cond1 := 0;cond2 := 0;cond3 := 0
	for i, c := range first {
		if i == 0 && c == 'i' || c == 'I' {
			cond1 = 1
		} else if c == 'a' || c == 'A' {
			cond2 = 1
		} else if i == len(first)-1 && c == 'n' || c == 'N' {
			cond3 = 1
		}
	}
	if cond1 == 1 && cond2 == 1 && cond3 == 1 {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
