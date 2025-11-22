package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i], &b[i])
	}
	c := make([]int, M)
	d := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&c[i], &d[i])
	}

	for i := 0; i < N; i++ {
		res := 0
		mini := int(2e9)
		for j := 0; j < M; j++ {
			dist := abs(a[i]-c[j]) + abs(b[i]-d[j])
			if mini > dist {
				mini = dist
				res = j + 1
			}
			mini = min(mini, dist)
		}
		fmt.Println(res)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}