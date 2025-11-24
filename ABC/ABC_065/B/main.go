package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
		a[i]--
	}

	res, tmp := 0, 0
	seen := make([]bool, N)
	for tmp != 1 {
		seen[tmp] = true
		tmp = a[tmp]
		res++
		
		if seen[tmp] {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(res)
}