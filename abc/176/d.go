package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w, ci, cj, di, dj int
	fmt.Scan(&h, &w)
	fmt.Scan(&ci, &cj)
	fmt.Scan(&di, &dj)

	s := make([][]string, h)

	for i := 0; i < h; i++ {
		var x string
		fmt.Scan(&x)
		s[i] = strings.Split(x, "")
	}

	// for i := 0; i < h; i++ {
	// 	for j := 0; j < w; j++ {
	// 		fmt.Printf("%s ", s[i][j])
	// 	}
	// 	fmt.Println("")
	// }

	wct := 0
	for i := 0; true; i++ {
		if ci > di && ci <= 0 && s[ci-1][cj] != "#" {
			ci--
		} else if ci < di && ci != h-1 && s[ci+1][cj] != "#" {
			ci++
		} else if cj > dj && cj <= 0 && s[ci][cj-1] != "#" {
			cj--
		} else if cj < dj && cj != w-1 && s[ci][cj+1] != "#" {
			cj++
		} else {

		}

		fmt.Printf("%d, %d\n", ci, cj)

		if ci == di && cj == dj {
			break
		}
	}
}
