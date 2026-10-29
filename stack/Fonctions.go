package stack

var Fonctions = map[string]func(a, b *[]int){
	"pa": Pa, "pb": Pb,
	"sa": Sa, "sb": Sb, "ss": Ss,
	"ra": Ra, "rb": Rb, "rr": Rr,
	"rra": Rra, "rrb": Rrb, "rrr": Rrr,
}
