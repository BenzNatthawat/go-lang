package loop

import "fmt"

func ForEachLoop() {
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
