package loop

import "fmt"

func WhileLoop() {
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}
