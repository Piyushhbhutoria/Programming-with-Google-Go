package main

import (
	"fmt"
)

func main() {
	var x int
	var arr []int
	for i := 0; i < 10; i++ {
		fmt.Println("Enter a no.")
		fmt.Scanln(&x)
		arr = append(arr, x)
	}
	sort := bubbleSort(arr)
	fmt.Println(sort)
}

func bubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
