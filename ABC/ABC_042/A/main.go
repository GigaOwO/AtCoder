package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])
	sort.Ints(a)

	if a[0] == 5 && a[1] == 5 && a[2] == 7 {
		fmt.Println("YES") 
	} else {
		fmt.Println("NO")
	}
} 