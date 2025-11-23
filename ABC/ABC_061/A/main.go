package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	if C >= A && C <= B {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}