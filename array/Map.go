package array

import "fmt"

func Map() {
	var mOne = make(map[string]string)
	var mTwo = map[string]string{}
	var mThree = map[string]string{
		"lastName": "YoYo",
	}

	fmt.Println("------------------- Map ----------------------")

	mOne["name"] = "John"
	mTwo["firstName"] = "Pawit"

	fmt.Printf("%v \n", mOne)
	fmt.Printf("%v \n", mTwo)
	fmt.Printf("%v \n", mThree)
}
