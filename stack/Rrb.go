package stack

func Rrb(a, b *[]int) {
	if len(*b) > 0 {
		last := (*b)[len(*b)-1]
		for i := len(*b) - 1; i > 0; i-- {
			(*b)[i] = (*b)[i-1]
		}
		(*b)[0] = last
	}
}
