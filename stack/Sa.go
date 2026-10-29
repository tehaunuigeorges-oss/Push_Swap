package stack

func Sa(a, b *[]int) {
	if len(*a) > 1 {
		(*a)[0], (*a)[1] = (*a)[1], (*a)[0]
	}
}
