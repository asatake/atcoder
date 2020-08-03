package main

import (
	"fmt"
)

func main() {
	ct := 0
	length := 0
	for i := 0; i < 3; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if tmp == 5 || tmp == 7 {
			ct++
		}
		length += tmp
	}
	if ct == 3 && length == 17 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
