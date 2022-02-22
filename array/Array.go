package array

import "fmt"

func Array() {
	var a [3]int

	a[0] = 5
	fmt.Println("------------------- Array ----------------------")

	fmt.Printf("%v \n", a[0])
	fmt.Printf("%v \n", a[1])
	fmt.Printf("%v \n", a[2])
}
