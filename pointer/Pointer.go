package pointer

import (
	"fmt"
)

func Pointer() {
	var s string = "Hello"
	var p *string

	fmt.Printf("variable s type %T value %v \n", p, p)

	p = &s
	fmt.Printf("variable s type %T value %v \n", p, p)
	fmt.Printf("variable s type %T value %v \n", *p, *p)

}
