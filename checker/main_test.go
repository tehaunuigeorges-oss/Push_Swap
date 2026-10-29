package main

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestChecker(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "checker.exe")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("compilation : %v\n%s", err, out)
	}
	cas := []struct {
		args, entree, sortie, erreur string
	}{
		{"3 2 1 0", "sa\nrra\npb\n", "KO\n", ""},
		{"3 2 one 0", "rra\npb\nsa\n", "", "Error\n"},
		{"3 2 1 0", "rra\npb\nsa\nrra\npa", "OK\n", ""},
		{"3 2 1", "rra\nxx\n", "", "Error\n"},
		{"1 2 3", "", "OK\n", ""},
		{"2 1", "pb\n", "KO\n", ""},
		{"2 1", "sa\r\n", "OK\n", ""},
		{"2 1", "\uFEFFsa\r\n", "OK\n", ""},
		{"1 2 1", "", "", "Error\n"},
		{"2 1", "sa \n", "", "Error\n"},
		{"", "sa\n", "", ""},
	}
	for _, c := range cas {
		var args []string
		if c.args != "" {
			args = []string{c.args}
		}
		cmd := exec.Command(bin, args...)
		cmd.Stdin = strings.NewReader(c.entree)
		var sortie, erreur strings.Builder
		cmd.Stdout, cmd.Stderr = &sortie, &erreur
		cmd.Run()
		if sortie.String() != c.sortie || erreur.String() != c.erreur {
			t.Errorf("checker %q < %q : sortie %q erreur %q, attendu %q %q", c.args, c.entree, sortie.String(), erreur.String(), c.sortie, c.erreur)
		}
	}
}
