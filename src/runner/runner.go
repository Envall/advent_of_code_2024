package main

import (
	"aoc24/dailys"
	"fmt"
)

func main() {
	// os package for command line args...

	dir := "/Users/jonathan_privat/Code/advent_of_code_2024/input_data/"
	file := "day4.txt"

	result := dailys.Day4_2(dir + file)

	fmt.Printf("Result: %d\n", result)
}
