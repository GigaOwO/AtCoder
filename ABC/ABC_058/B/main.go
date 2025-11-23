package main

import (
	"fmt"
)

func main() {
	var O, E string
	fmt.Scan(&O, &E)

	for i := 0; i < len(E); i++ {
		fmt.Printf("%c%c", O[i], E[i])
	}
	if len(O) > len(E) {
		fmt.Printf("%c", O[len(O) - 1])
	}
}