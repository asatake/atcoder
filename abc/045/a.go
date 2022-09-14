package abc

import "fmt"

func main() {
	var a, b, h int64
	fmt.Scan(&a, &b, &h)

	fmt.Print((a + b) * h / 2)
}
