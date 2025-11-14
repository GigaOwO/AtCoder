package main

import "fmt"
 
func main() {
	var N, K, D int
	fmt.Scan(&N, &K)
	a := make([]bool, 10)
	for i := 0; i < K; i++ {
		fmt.Scan(&D)
		a[D] = true
	}
	for check(a, N) {
		N++
	}
	fmt.Println(N)
}

func check(a []bool, N int) bool {
	for N > 0 {
		if a[N%10] {
			return true
		}
		N /= 10
	}
	return false
}