package smoke

func Total(xs []int) int {
	n := 0
	for _, x := range xs {
		n += x
	}
	return n
}
