package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseArgs(args []string) ([]int, error) {
	var morceaux []string
	for _, arg := range args {
		morceaux = append(morceaux, strings.Fields(arg)...)
	}

	seen := make(map[int]bool)
	var nums []int
	for _, morceau := range morceaux {
		n, err := strconv.Atoi(morceau)
		if err != nil {
			return nil, err
		}
		if seen[n] {
			return nil, fmt.Errorf("doublon : %d", n)
		}
		seen[n] = true

		nums = append(nums, n)
	}
	return nums, nil
}
