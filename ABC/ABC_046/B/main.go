package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	res := K
	for i := 0; i < N - 1; i++ {
		res *= K - 1
	}
	fmt.Println(res)
}