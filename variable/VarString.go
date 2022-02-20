package variable

import (
	"fmt"
)

func VarString() {
	fmt.Println("------------------- string ----------------------")

	var sOne string

	fmt.Printf("variable s_one type %T value %v \n", sOne, sOne)
	sOne = "text"
	fmt.Printf("variable s_one type %T value %v \n", sOne, sOne)

	sTwo := "Hello"
	fmt.Printf("variable sTwo type %T value %v \n", sTwo, sTwo)
	sTwo = "World"
	fmt.Printf("variable sTwo type %T value %v \n", sTwo, sTwo)

	fmt.Println("")
}
