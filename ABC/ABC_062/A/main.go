package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if group(x) == group(y) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func group(x int) int {
	switch x {
	case 2:
		return 0
	case 1, 3, 5, 7, 8, 10, 12:
		return 1
	default:
		return 2
	}
}