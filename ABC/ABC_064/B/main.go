package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	res := 0
	for i := 0; i < N - 1; i++ {
		res += a[i + 1] - a[i] 
	}
	fmt.Println(res)
}