package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]string)
	var x string
	var y string
	fmt.Println("Enter your name")
	fmt.Scanln(&x)
	fmt.Println("Enter your address")
	fmt.Scanln(&y)
	m["name"] = x
	m["address"] = y
	json, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}
