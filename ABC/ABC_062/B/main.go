package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)
	a := make([]string, H)
	for i := range a {
		fmt.Scan(a[i])
	}
	
	for i := 0; i < W + 2; i++ {
		fmt.Print("#")
	}
	fmt.Println()
	for i := 0; i < H; i++ {
		fmt.Print("#", a[i], "#", "\n")
	}
	for i := 0; i < W + 2; i++ {
		fmt.Print("#")
	}
	fmt.Println()
}