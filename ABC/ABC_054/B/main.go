package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	A, B := make([]string, N), make([]string, M)
	for i := range A {
		fmt.Scan(&A[i])
	}
	for i := range B {
		fmt.Scan(&B[i])
	}

	res := "No"
	for i := 0; i < N - M + 1; i++ {
		for j := 0; j < N - M + 1; j++ {
			cnt := 0
			for y := 0; y < M; y++ {
				for x := 0; x < M; x++ {
					if A[i + y][j + x] == B[y][x] {
						cnt++
					}
				}
			}
			if cnt == M * M {
				res = "Yes"
			}
		}
	}

	fmt.Println(res)
}
