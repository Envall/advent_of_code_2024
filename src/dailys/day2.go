package dailys

import (
	"aoc24/functions"
	"fmt"
	"math"
)

// Checks if a report breaks any of the two rules (is "unsafe").
// skips checking number at index <skip>
// Returns true if unsafe, and the index where the report became unsafe.
func check_if_unsafe_report(report []int, skip int) (bool, int) {
	prev_val := 0
	prev_dir := 0
	unsafe := false
	i_unsafe := -1

	for i, val := range report {
		if i == 0 || (skip == 0 && i == 1) {
			prev_val = val
			prev_dir = 0
			continue
		} else if i == skip {
			continue
		}

		var dir int
		diff := math.Abs(float64(val - prev_val))
		if diff > 0 {
			dir = (val - prev_val) / int(diff) // 1 or -1
		} else {
			dir = 0
		}

		// Check if safe
		if diff > 3 || diff < 1 {
			unsafe = true
			i_unsafe = i
			break
		} else if dir*prev_dir == -1 {
			unsafe = true
			i_unsafe = i
			break
		}

		prev_val = val
		prev_dir = dir
		if dir != 0 {
			prev_dir = dir
			// If the sequence is e.g. [1 2 2 1],
			// we want to keep dir=1 for i=[1:2]
		}
	}

	return unsafe, i_unsafe
}

func Day2_1(path string) int {
	fmt.Println("Running Day2 part 1")

	data := functions.Readf_lines_as_int_slices(path)

	//fmt.Println(data)

	n_safe_reports := 0

	for _, report := range data {
		unsafe, _ := check_if_unsafe_report(report, -1)
		if !unsafe {
			n_safe_reports++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", n_safe_reports)
	return n_safe_reports
}

func Day2_2(path string) int {
	fmt.Println("Running Day2 part 2")

	data := functions.Readf_lines_as_int_slices(path)

	n_safe_reports := 0

	for _, report := range data {
		unsafe, _ := check_if_unsafe_report(report, -1)
		if unsafe {
			// Check if report becomes safe if we remove the index where
			// report became unsafe, or the index before.
			//unsafe2, _ := check_if_unsafe_report(report, unsafe_i-1)
			//unsafe3, _ := check_if_unsafe_report(report, unsafe_i)
			//if !unsafe2 || !unsafe3 {
			//	n_safe_reports++
			//}

			// Try removing each element in the report and see if that
			// makes it safe...
			//
			// This worked but why? Compare result with above method?
			for rm := range len(report) {
				unsafe, _ := check_if_unsafe_report(report, rm)
				if !unsafe {
					n_safe_reports++
					break
				}
			}
		} else {
			n_safe_reports++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", n_safe_reports)
	return n_safe_reports
}
