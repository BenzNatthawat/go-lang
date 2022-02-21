package conditions

import "fmt"

func IfElse(name string) {
	if name != "" {
		fmt.Printf("Hello %s\n", name)
	} else {
		fmt.Printf("Hello World\n")
	}
}
