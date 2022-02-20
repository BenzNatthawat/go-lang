package variable

import (
	"fmt"
)

func VarBoolean() {
	fmt.Println("------------------- boolean ----------------------")

	var bOne bool
	fmt.Printf("variable s type %T value %v \n", bOne, bOne)
	bOne = true
	fmt.Printf("variable s type %T value %v \n", bOne, bOne)

	bTwo := false
	fmt.Printf("variable s type %T value %v \n", bTwo, bTwo)
	bTwo = true
	fmt.Printf("variable s type %T value %v \n", bTwo, bTwo)

	fmt.Println("")
}
