package variable

import "fmt"

func VarInt() {
	fmt.Println("------------------- int ----------------------")

	var iOne int
	fmt.Printf("variable iOne type %T value %v \n", iOne, iOne)
	iOne = 10
	fmt.Printf("variable iOne type %T value %v \n", iOne, iOne)

	iTwo := 0
	fmt.Printf("variable iTwo type %T value %v \n", iTwo, iTwo)
	iTwo = 10
	fmt.Printf("variable iTwo type %T value %v \n", iTwo, iTwo)

	fmt.Println("")
}
