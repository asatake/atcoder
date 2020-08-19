package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	k := 0
	for i := 1; i <= n; i++ {
		k += i
	}
	fmt.Println(k)
}
