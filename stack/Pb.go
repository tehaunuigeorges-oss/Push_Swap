package stack

func Pb(a, b *[]int) {
	if len(*a) > 0 {
		*b = append([]int{(*a)[0]}, (*b)...)
		*a = (*a)[1:]
	}
}
