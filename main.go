package main

import (
	"fmt"

	"./functions"
)

func main() {
	functions.NoReturn("Benz", "Sing")

	var resultOneReturn = functions.OneReturn(10, 20)
	fmt.Printf("result %v \n", resultOneReturn)

	var resultOne, resultTwo = functions.ManyReturn(10, 20)
	fmt.Printf("result plus one: %v \n", resultOne)
	fmt.Printf("result multiply two: %v \n", resultTwo)
}
