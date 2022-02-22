package array

import "fmt"

func AppendArray() {
	var sl = make([]int, 4)

	fmt.Println("------------------- AppendArray ----------------------")
	sl[0] = 5
	sl[2] = 60
	sl = append(sl, 20)

	fmt.Printf("%v \n", sl[0])
	fmt.Printf("%v \n", sl[1])
	fmt.Printf("%v \n", sl[2])
	fmt.Printf("%v \n", sl[3])
	fmt.Printf("%v \n", sl)
}
