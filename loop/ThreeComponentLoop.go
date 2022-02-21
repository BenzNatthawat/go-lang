package loop

import "fmt"

func ThreeComponentLoop() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
}
