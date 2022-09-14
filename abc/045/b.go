package abc

import "fmt"

func main() {
	var sa, sb, sc string
	fmt.Scan(&sa, &sb, &sc)

	turn := "a"

	for {
		switch turn {
		case "a":
			if len(sa) > 0 {
				turn, sa = t(turn, sa)
			} else {
				fmt.Printf("A")
				return
			}
		case "b":
			if len(sb) > 0 {
				turn, sb = t(turn, sb)
			} else {
				fmt.Printf("B")
				return
			}
		case "c":
			if len(sc) > 0 {
				turn, sc = t(turn, sc)
			} else {
				fmt.Printf("C")
				return
			}
		}
	}
}

func t(turn string, s string) (string, string) {
	turn = s[:1]
	s = s[1:]
	return turn, s
}
