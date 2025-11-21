package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N)
	T := make([]int, N)
	for i := range T {
		fmt.Scan(&T[i])
	}
	fmt.Scan(&M)
	P, X := make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&P[i], &X[i])
	}
	
	for i := 0; i < M; i++ {
		res := 0
		for j := 0; j < N; j++ {
			if j == P[i] - 1 {
				res += X[i]
			} else {
				res += T[j]
			}
		}
		fmt.Println(res)
	}
}