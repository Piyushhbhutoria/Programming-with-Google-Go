package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s []int
	var x string
	fmt.Scanln(&x)
	for x != "X" {
		a, _ := strconv.Atoi(x)
		s = append(s, a)
		sort.Ints(s)
		fmt.Println(s)
		fmt.Scanln(&x)
	}
}
