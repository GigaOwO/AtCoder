package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)
	n := len(S)
	res := 0
	
	for bit := 0; bit < 1<<(n-1); bit++ {
		tmp := ""
		for i := 0; i < n; i++ {
			tmp += string(S[i])
			if bit&(1<<i) == 0 {
				t, _ := strconv.Atoi(tmp)
				res += t
				tmp = ""
			}
		}
	}
	fmt.Println(res)
}