package array

import "fmt"

func Slice() {
	var sl = make([]int, 4)
	fmt.Println("------------------- Slice ----------------------")

	sl[0] = 5
	sl[2] = 9
	fmt.Printf("%v \n", sl[0])
	fmt.Printf("%v \n", sl[1])
	fmt.Printf("%v \n", sl[2])
	fmt.Printf("%v \n", sl[3])
}
