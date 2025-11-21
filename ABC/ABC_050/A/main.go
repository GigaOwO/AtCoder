package main

import (
	"fmt"
)

func main() {
	var A, B int
	var op string
	fmt.Scan(&A, &op, &B)

	switch op {
		case "+":
			fmt.Println(A + B)
		case "-":
			fmt.Println(A - B)
		}
}