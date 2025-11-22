package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if A == 1 {
		A += 13
	}
	if B == 1 {
		B += 13
	}

	if A > B {
		fmt.Println("Alice")
	} else if A < B {
		fmt.Println("Bob")
	} else {
		fmt.Println("Draw")
	}
}