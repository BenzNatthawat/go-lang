package array

func VariadicFunc(n ...int) []int {
	for i := 0; i < len(n); i++ {
		n[i] = n[i] * 2
	}
	return n
}
