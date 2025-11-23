package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	if C % gcd(A, B) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x % y)
}