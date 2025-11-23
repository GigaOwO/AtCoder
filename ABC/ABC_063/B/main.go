package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)

	m := map[byte]struct{}{}
	for i := range S {
		m[S[i]] = struct{}{}
	}
	if len(S) == len(m) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}