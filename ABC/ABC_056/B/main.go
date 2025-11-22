package main

import "fmt"

func main() {
	var W, a, b int
	fmt.Scan(&W, &a, &b)
	if a > b {
		a, b = b, a
	}
	
	if a + W >= b {
		fmt.Println(0)
	} else {
		fmt.Println(b - (a + W))
	}
}