package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	res := ""

	for _, c := range s {
		if c == '0' || c == '1' {
			res += string(c)
		} else if len(res) > 0 {
			res = res[:len(res) - 1]
		}
	}

	fmt.Println(res)
}