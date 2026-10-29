package stack

func Rra(a, b *[]int) {
	if len(*a) > 0 {
		last := (*a)[len(*a)-1]
		for i := len(*a) - 1; i > 0; i-- {
			(*a)[i] = (*a)[i-1]
		}
		(*a)[0] = last
	}
}
