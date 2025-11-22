package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	res := 1
	for i := 1; i <= N; i++ {
		res = res * i % int(1e9+7)
	}
	fmt.Println(res)
}