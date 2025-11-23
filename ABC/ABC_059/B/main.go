package main

import "fmt"

func main() {
	var A, B string
	fmt.Scan(&A, &B)

	if len(A) > len(B) {
		fmt.Println("GREATER")
	} else if len(A) < len(B) {
		fmt.Println("LESS")
	} else if A > B {
		fmt.Println("GREATER")
	} else if A < B {
		fmt.Println("LESS")
	} else {
		fmt.Println("EQUAL")
	}
}