package constants

import (
	"fmt"
)

func ConstDefault() {
	fmt.Println("------------------- boolean ----------------------")

	const defaultInt = 1
	const defaultString = "Go"
	const defaultBoolean = true

	fmt.Printf("defaultInt %T, %v \n", defaultInt, defaultInt)
	fmt.Printf("defaultInt %T, %v \n", defaultString, defaultString)
	fmt.Printf("defaultInt %T, %v \n", defaultBoolean, defaultBoolean)
}
