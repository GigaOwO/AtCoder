package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	res := 0
	count := 0
	for i := 0; i < N; i++ {
		switch S[i] {
			case 'I':
				count++
				res = max(res, count)
			case 'D':
				count--
		}
	}
	fmt.Println(res)
}