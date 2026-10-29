package stack

func Ra(a, b *[]int) {
	if len(*a) > 0 {
		first := (*a)[0]
		for i := 0; i < len(*a)-1; i++ {
			(*a)[i] = (*a)[i+1]
		}
		(*a)[len(*a)-1] = first
	}
}
