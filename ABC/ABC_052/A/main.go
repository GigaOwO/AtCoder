package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)
	if A * B >= C * D {
		fmt.Println(A * B)
	} else {
		fmt.Println(C * D)
	}
}