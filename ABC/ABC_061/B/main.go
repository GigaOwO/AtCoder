package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	a, b := make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&a[i], &b[i])
	}

	res := make([]int, N)
	for i := 0; i < M; i++ {
		res[a[i] - 1]++
		res[b[i] - 1]++
	}
	for i := range res {
		fmt.Println(res[i])
	}
}