package array

import "fmt"

func AppendArray() {
	var sl []int

	sl = make([]int, 4)

	sl[0] = 5
	fmt.Println("------------------- AppendArray ----------------------")

	sl = append(sl, 20)
	fmt.Printf("%v \n", sl[0])
	fmt.Printf("%v \n", sl[1])
	fmt.Printf("%v \n", sl[2])
	fmt.Printf("%v \n", sl[3])
	fmt.Printf("%v \n", sl[4])
}
