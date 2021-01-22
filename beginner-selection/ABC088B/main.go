package ABC088B

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, 2)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	alice, bob := 0, 0
	for i, c := range a {
		if i%2 == 0 {
			alice += c
		} else {
			bob += c
		}
	}

	if alice >= bob {
		fmt.Println(alice - bob)
	} else {
		fmt.Println(bob - alice)
	}
}
