package loop

import "fmt"

func ExitLoop() {
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
