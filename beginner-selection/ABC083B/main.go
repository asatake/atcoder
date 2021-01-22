package main

import (
	"fmt"
)

func sum(slice []int) int {
	sum := 0
	for _, s := range slice {
		sum += s
	}
	return sum
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	total := 0
	for i := 1; i <= n; i++ {
		s := []int{}
		x := i
		for {
			s = append(s, x%10)
			if x/10 == 0 {
				break
			}
			x /= 10
		}
		j := sum(s)
		if j >= a && j <= b {
			total += i
		}

	}
	fmt.Println(total)
}
