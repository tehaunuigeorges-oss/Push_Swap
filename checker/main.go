package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"

	"github.com/tehaunuigeorges-oss/push-swap/parse"
	"github.com/tehaunuigeorges-oss/push-swap/stack"
)

func main() {
	args := os.Args[1:]
	auto := len(args) > 0 && args[0] == "-auto"
	if auto {
		args = args[1:]
	}
	a, err := parse.ParseArgs(args)
	if err != nil {
		erreur()
	}
	if len(a) == 0 {
		return
	}
	var instructions []func(a, b *[]int)
	var source io.Reader = os.Stdin
	if auto {
		cmd := exec.Command("go", append([]string{"run", "./push-swap"}, args...)...)
		cmd.Stderr = os.Stderr
		sortie, err := cmd.Output()
		if err != nil {
			erreur()
		}
		os.Stdout.Write(sortie)
		source = bytes.NewReader(sortie)
	}
	lecteur := bufio.NewScanner(source)
	for lecteur.Scan() {
		ligne := strings.TrimSuffix(strings.TrimPrefix(lecteur.Text(), "\uFEFF"), "\r")
		if ligne == "" {
			continue
		}
		f, ok := stack.Fonctions[ligne]
		if !ok {
			erreur()
		}
		instructions = append(instructions, f)
	}
	if lecteur.Err() != nil {
		erreur()
	}
	var b []int
	for _, f := range instructions {
		f(&a, &b)
	}
	if len(b) == 0 && slices.IsSorted(a) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func erreur() {
	fmt.Fprintln(os.Stderr, "Error")
	os.Exit(1)
}
