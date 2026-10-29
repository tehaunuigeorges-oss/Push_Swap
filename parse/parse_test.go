package parse

import (
	"slices"
	"testing"
)

func TestParseArgs(t *testing.T) {
	valides := [][]string{{"2", "1", "3"}, {"2 1 3"}, {"  2   1 ", "3"}}
	for _, args := range valides {
		nums, err := ParseArgs(args)
		if err != nil || !slices.Equal(nums, []int{2, 1, 3}) {
			t.Errorf("ParseArgs(%q) = %v, %v ; attendu [2 1 3]", args, nums, err)
		}
	}
	erreurs := [][]string{{"one"}, {"2.5"}, {"0 one 2 3"}, {"2 2"}, {"2", "1", "2"}, {"99999999999999999999"}}
	for _, args := range erreurs {
		if _, err := ParseArgs(args); err == nil {
			t.Errorf("ParseArgs(%q) devrait renvoyer une erreur", args)
		}
	}
}

func TestParseArgsVide(t *testing.T) {
	for _, args := range [][]string{nil, {""}, {"   "}} {
		if nums, err := ParseArgs(args); err != nil || len(nums) != 0 {
			t.Errorf("ParseArgs(%q) = %v, %v ; attendu rien", args, nums, err)
		}
	}
}

func TestParseArgsNegatifs(t *testing.T) {
	nums, err := ParseArgs([]string{"-3 0 +2"})
	if err != nil || !slices.Equal(nums, []int{-3, 0, 2}) {
		t.Errorf("ParseArgs = %v, %v ; attendu [-3 0 2]", nums, err)
	}
}
