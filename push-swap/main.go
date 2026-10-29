package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/tehaunuigeorges-oss/push-swap/parse"
	"github.com/tehaunuigeorges-oss/push-swap/stack"
)

type pileState struct {
	a, b        []int
	instruction []string
}

func main() {
	nums, err := parse.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	if len(nums) == 0 {
		return
	}
	instruction := solve(nums)
	if len(instruction) > 0 {
		fmt.Println(strings.Join(instruction, "\n"))
	}
	fmt.Fprintf(os.Stderr, "Nombre d'instructions : %d\n", len(instruction))
}

func solve(nums []int) []string {
	p := &pileState{a: nums}
	switch {
	case EstRangee(p):
	case len(p.a) == 2:
		p.faire("sa")
	case len(p.a) == 3:
		trier3(p)
	case len(p.a) <= 7:
		trier5(p)
	case len(p.a) <= 12:
		trierMoyen(p)
	default:
		trierGrand(p)
	}
	return p.instruction
}

func trier3(p *pileState) {
	switch max(p.a[0], p.a[1], p.a[2]) {
	case p.a[0]:
		p.faire("ra")
	case p.a[1]:
		p.faire("rra")
	}
	if p.a[0] > p.a[1] {
		p.faire("sa")
	}
}

func trier5(p *pileState) {
	type etat struct {
		a, b   []int
		chemin []string
	}
	noms := []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}
	vus := map[string]bool{fmt.Sprint(p.a, p.b): true}
	file := []etat{{a: p.a}}
	for len(file) > 0 {
		e := file[0]
		file = file[1:]
		for _, nom := range noms {
			a, b := slices.Clone(e.a), slices.Clone(e.b)
			stack.Fonctions[nom](&a, &b)
			cle := fmt.Sprint(a, b)
			if vus[cle] {
				continue
			}
			vus[cle] = true
			chemin := append(slices.Clone(e.chemin), nom)
			if len(b) == 0 && slices.IsSorted(a) {
				for _, op := range chemin {
					p.faire(op)
				}
				return
			}
			file = append(file, etat{a, b, chemin})
		}
	}
}

func trierMoyen(p *pileState) {
	type etat struct {
		a, b   []int
		chemin []string
		note   int
	}
	const largeur = 2000
	noms := []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}
	tri := slices.Sorted(slices.Values(p.a))
	depart := make([]int, len(p.a))
	for i, v := range p.a {
		depart[i], _ = slices.BinarySearch(tri, v)
	}
	cle := func(a, b []int) string {
		k := make([]byte, 0, len(a)+len(b)+1)
		for _, v := range a {
			k = append(k, byte(v))
		}
		k = append(k, 255)
		for _, v := range b {
			k = append(k, byte(v))
		}
		return string(k)
	}
	noter := func(a, b []int) int {
		n := 2 * len(b)
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1]+1 {
				n++
			}
		}
		for i := 1; i < len(b); i++ {
			if b[i] != b[i-1]-1 {
				n++
			}
		}
		return n
	}
	grand := &pileState{a: slices.Clone(p.a)}
	trierGrand(grand)
	meilleur := grand.instruction
	vus := map[string]bool{cle(depart, nil): true}
	niveau := []etat{{a: depart}}
	for len(niveau) > 0 && len(niveau[0].chemin)+1 < len(meilleur) {
		var suivant []etat
		for _, e := range niveau {
			for _, nom := range noms {
				a, b := slices.Clone(e.a), slices.Clone(e.b)
				stack.Fonctions[nom](&a, &b)
				k := cle(a, b)
				if vus[k] {
					continue
				}
				vus[k] = true
				chemin := append(slices.Clone(e.chemin), nom)
				if len(b) == 0 && slices.IsSorted(a) {
					meilleur = chemin
					break
				}
				suivant = append(suivant, etat{a, b, chemin, noter(a, b)})
			}
			if len(meilleur) == len(e.chemin)+1 {
				break
			}
		}
		slices.SortStableFunc(suivant, func(x, y etat) int { return cmp.Compare(x.note, y.note) })
		niveau = suivant[:min(len(suivant), largeur)]
	}
	for _, op := range meilleur {
		p.faire(op)
	}
}

func trierGrand(p *pileState) {
	tri := slices.Sorted(slices.Values(p.a))
	pivot := tri[len(tri)/2]
	rb := false
	for len(p.a) > 3 {
		deplacer(p, true, rb)
		rb = p.b[0] < pivot
	}
	trier3(p)
	for len(p.b) > 0 {
		deplacer(p, false, false)
	}
	i := slices.Index(p.a, slices.Min(p.a))
	if i > len(p.a)/2 {
		i -= len(p.a)
	}
	for ; i > 0; i-- {
		p.faire("ra")
	}
	for ; i < 0; i++ {
		p.faire("rra")
	}
}

func deplacer(p *pileState, versB bool, rb bool) {
	src, dst, rs, rd, push := p.a, p.b, "ra", "rb", "pb"
	if !versB {
		src, dst, rs, rd, push = p.b, p.a, "rb", "ra", "pa"
	}
	cout := func(x, y int) int {
		if (x < 0) == (y < 0) {
			return max(x, y, -x, -y)
		}
		return max(x, -x) + max(y, -y)
	}
	bi, bj, best := 0, 0, -1
	for i, v := range src {
		j := 0
		if versB {
			if rb && len(dst) > 1 {
				j = 1
			}
		} else {
			j = slices.Index(dst, slices.Min(dst))
			for k, w := range dst {
				if w > v && (dst[j] < v || w < dst[j]) {
					j = k
				}
			}
		}
		ls, ld := len(src), len(dst)
		for _, o := range [][2]int{{i, j}, {i - ls, j - ld}, {i, j - ld}, {i - ls, j}} {
			if c := cout(o[0], o[1]); best < 0 || c < best {
				bi, bj, best = o[0], o[1], c
			}
		}
	}
	for ; bi > 0 && bj > 0; bi, bj = bi-1, bj-1 {
		p.faire("rr")
	}
	for ; bi < 0 && bj < 0; bi, bj = bi+1, bj+1 {
		p.faire("rrr")
	}
	for ; bi > 0; bi-- {
		p.faire(rs)
	}
	for ; bi < 0; bi++ {
		p.faire("r" + rs)
	}
	for ; bj > 0; bj-- {
		p.faire(rd)
	}
	for ; bj < 0; bj++ {
		p.faire("r" + rd)
	}
	p.faire(push)
}

func (p *pileState) faire(op string) {
	stack.Fonctions[op](&p.a, &p.b)
	p.instruction = append(p.instruction, op)
}

func EstRangee(p *pileState) bool {
	return slices.IsSorted(p.a)
}
