package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	const INF = int64(1 << 60)
	res := INF

	for x := -100; x <= 100; x++ {
		var tmp int64 = 0
		for i := 0; i < N; i++ {
			d := int64(x - a[i])
			tmp += d * d
		}
		if tmp < res {
			res = tmp
		}
	}
	fmt.Println(res)
}
