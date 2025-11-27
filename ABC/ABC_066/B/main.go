package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	for {
		S = S[:len(S) - 2]
		if len(S) % 2 != 0 {
			continue
		}
		if S[:len(S) / 2] == S[len(S) / 2:] {
			break
		}
	}
	fmt.Println(len(S))
}