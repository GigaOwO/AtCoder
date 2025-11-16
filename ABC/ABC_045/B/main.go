package main

import "fmt"

func main() {
	S := make([]string, 3)
	for i := range S {
		fmt.Scan(&S[i])
	}
	
	i := 0
	for len(S[i]) > 0 {
		j := S[i][0] - 'a'
		S[i] = S[i][1:]
		i = int(j)
	}

	fmt.Println(string('A' + byte(i)))
}