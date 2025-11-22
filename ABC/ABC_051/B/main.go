package main

import "fmt"

func main() {
	var K, S int
	fmt.Scan(&K, &S)

	res := 0
	for x := 0; x <= K; x++ {
		for y := 0; y <= K; y++ {
			if 0 <= S - x - y && S - x - y <= K {
				res++
			}
		}
	}

	fmt.Println(res)
}