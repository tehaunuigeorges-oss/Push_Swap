package main

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestProgramme(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "push-swap.exe")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("compilation : %v\n%s", err, out)
	}
	cas := []struct {
		args           []string
		sortie, erreur string
	}{
		{nil, "", ""},
		{[]string{"0 one 2 3"}, "", "Error\n"},
		{[]string{"1 2 1"}, "", "Error\n"},
		{[]string{"2 1"}, "sa\n", "Nombre d'instructions : 1\n"},
		{[]string{"1 2 3"}, "", "Nombre d'instructions : 0\n"},
	}
	for _, c := range cas {
		cmd := exec.Command(bin, c.args...)
		var sortie, erreur strings.Builder
		cmd.Stdout, cmd.Stderr = &sortie, &erreur
		cmd.Run()
		if sortie.String() != c.sortie || erreur.String() != c.erreur {
			t.Errorf("push-swap %q : sortie %q erreur %q, attendu %q %q", c.args, sortie.String(), erreur.String(), c.sortie, c.erreur)
		}
	}
}
