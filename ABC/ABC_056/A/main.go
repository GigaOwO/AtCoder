package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	if a == "H" {
		if b == "H" {
			fmt.Println(a)
		} else {
			fmt.Println(b)
		}
	} else {
		if b == "H" {
			fmt.Println(a)
		} else {
			fmt.Println("H")
		}
	}
}