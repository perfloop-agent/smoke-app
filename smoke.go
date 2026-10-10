package smoke

func Total(xs []int) int {
	n := 0
	for i := range xs {
		n += xs[i]
	}
	return n
}
