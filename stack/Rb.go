package stack

func Rb(a, b *[]int) {
	if len(*b) > 0 {
		first := (*b)[0]
		for i := 0; i < len(*b)-1; i++ {
			(*b)[i] = (*b)[i+1]
		}
		(*b)[len(*b)-1] = first
	}
}
