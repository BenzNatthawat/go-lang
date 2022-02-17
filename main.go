package main

import "fmt"

func main() {
	var i int
	fmt.Printf("variable i type %T value %v \n", i, i)
	i = 10
	fmt.Printf("variable i type %T value %v \n", i, i)

	var s string
	fmt.Printf("variable s type %T value %v \n", s, s)
	s = "text"
	fmt.Printf("variable s type %T value %v \n", s, s)

	var b bool
	fmt.Printf("variable s type %T value %v \n", b, b)
	b = true
	fmt.Printf("variable s type %T value %v \n", b, b)

	var st *string
	fmt.Printf("variable s type %T value %v \n", st, st)
}
