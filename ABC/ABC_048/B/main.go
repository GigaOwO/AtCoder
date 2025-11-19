package main

import "fmt"

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)
	
	lower := -1
	if a != 0 {
		lower = (a - 1) / x
	}
	upper := b / x
	fmt.Println(upper - lower)
}