package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ls := []int{}
	for _, s := range strings.Split(sc.Text(), " ") {
		num, _ := strconv.Atoi(s)
		ls = append(ls, num)
	}

	length := len(ls)
	ct := 0

	for i, x := range ls {
		if i >= length - 2 {
			break
		}
		for j, y := range ls[i+1:] {
			if i + j + 1 >= length - 1 {
				break
			}
			if x == y {
				continue
			}
			for _, z := range ls[i+j+2:] {
				if y == z || z == x {
					continue
				}
				if x + y > z && y + z > x && z + x > y {
					ct++
				}
			}
		}
	}
	fmt.Println(ct)
}
