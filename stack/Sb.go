package stack

func Sb(a, b *[]int) {
	if len(*b) > 1 {
		(*b)[0], (*b)[1] = (*b)[1], (*b)[0]
	}
}
