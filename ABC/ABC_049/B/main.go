package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	s := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scan(&s[i])
	}
	for i := 0; i < H; i++ {
		fmt.Println(s[i])
		fmt.Println(s[i])
	}
}