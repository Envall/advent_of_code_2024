package dailys

import (
	"aoc24/functions"
	"fmt"
	"math"
	"sort"
)

func Day1_1(path string) {
	c1, c2 := functions.Readf_2cols_int(path)

	sort.Slice(c1, func(i, j int) bool {
		return c1[i] < c1[j]
	})
	sort.Slice(c2, func(i, j int) bool {
		return c2[i] < c2[j]
	})

	if len(c1) != len(c2) {
		panic("Bad input dat! Different length arrays..")
	}

	distance := 0
	for i := range len(c1) {
		distance += int(math.Abs(float64(c1[i] - c2[i])))
	}

	fmt.Printf("Part 1 distance: %d\n", distance)
}

func Day1_2(path string) {
	c1, c2 := functions.Readf_2cols_int(path)
	dist := 0
	for i := range len(c1) {
		dist += c1[i] * functions.Count_in_slice(c1[i], c2)
	}
	fmt.Printf("Part 2 distance: %d\n", dist)
}
