package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)
	fmt.Println(strings.ToUpper(s1[:1] +  s2[:1] + s3[:1]))
}