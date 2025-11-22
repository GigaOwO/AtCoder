package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	A, Z := 0, 0
	A = strings.Index(s, "A")
	Z = strings.LastIndex(s, "Z")
	fmt.Println(Z - A + 1)
}