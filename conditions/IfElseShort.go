package conditions

import (
	"fmt"
	"strconv"
)

func IfElseShort() {
	if n, err := strconv.Atoi("5b"); err != nil {
		fmt.Printf("Nan: %v", err)
	} else {
		fmt.Printf("%v %T", n, n)
	}
}
