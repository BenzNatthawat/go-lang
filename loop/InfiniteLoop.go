package loop

import "fmt"

func InfiniteLoop() {
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
	}
}
