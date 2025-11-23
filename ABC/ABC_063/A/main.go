package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if A + B >= 10 {
		fmt.Println("error")
	} else {
		fmt.Println(A + B)
	}
}