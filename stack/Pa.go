package stack

func Pa(a, b *[]int) {
	if len(*b) > 0 {
		*a = append([]int{(*b)[0]}, (*a)...)
		*b = (*b)[1:]
	}
}
