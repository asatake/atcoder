package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	length := len(s)

	for i := 0; i < length - 2; i++ {
		if s[i] == s[i+2] {
			fmt.Printf("%d %d\n", i + 1, i + 3)
			return
		}
		if s[i] == s[i+1] {
			fmt.Printf("%d %d\n", i + 1, i + 2)
			return
		}
	}
	if s[length - 2] == s[length - 1] {
		fmt.Printf("%d %d", length - 1, length)
	} else {
		fmt.Println("-1 -1")
	}
}
