package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/tehaunuigeorges-oss/push-swap/stack"
)

func rejouer(t *testing.T, nums []int, instruction []string) bool {
	t.Helper()
	p := &pileState{a: slices.Clone(nums)}
	for _, op := range instruction {
		f, ok := stack.Fonctions[op]
		if !ok {
			t.Fatalf("instruction inconnue %q", op)
		}
		f(&p.a, &p.b)
	}
	return len(p.b) == 0 && EstRangee(p)
}

func permutations(n int) [][]int {
	if n == 0 {
		return [][]int{{}}
	}
	var res [][]int
	for _, p := range permutations(n - 1) {
		for i := 0; i <= len(p); i++ {
			res = append(res, slices.Insert(slices.Clone(p), i, n-1))
		}
	}
	return res
}

func TestPetitesPiles(t *testing.T) {
	pireAttendu := map[int]int{1: 0, 2: 1, 3: 2, 4: 5, 5: 8}
	for n, attendu := range pireAttendu {
		pire := 0
		for _, nums := range permutations(n) {
			instruction := solve(slices.Clone(nums))
			if !rejouer(t, nums, instruction) {
				t.Fatalf("%v n'est pas rangee par %v", nums, instruction)
			}
			pire = max(pire, len(instruction))
		}
		if pire != attendu {
			t.Errorf("n=%d : pire cas %d instructions, attendu %d", n, pire, attendu)
		}
	}
}

func TestGrandesPiles(t *testing.T) {
	limites := map[int]int{6: 12, 7: 14, 8: 18, 10: 26, 12: 42, 100: 700, 500: 5500}
	r := rand.New(rand.NewSource(42))
	for n, limite := range limites {
		for i := 0; i < 10; i++ {
			nums := r.Perm(n)
			instruction := solve(slices.Clone(nums))
			if !rejouer(t, nums, instruction) {
				t.Fatalf("%v n'est pas rangee", nums)
			}
			if len(instruction) > limite {
				t.Errorf("n=%d : %d instructions, limite %d", n, len(instruction), limite)
			}
		}
	}
}

func TestExemplesSujet(t *testing.T) {
	cas := []struct {
		nums []int
		max  int
	}{
		{[]int{4, 67, 3, 87, 23}, 6},
		{[]int{2, 1, 3, 6, 5, 8}, 7},
	}
	for _, c := range cas {
		instruction := solve(slices.Clone(c.nums))
		if !rejouer(t, c.nums, instruction) || len(instruction) > c.max {
			t.Errorf("%v : %d instructions %v, le sujet en attend %d", c.nums, len(instruction), instruction, c.max)
		}
	}
}
