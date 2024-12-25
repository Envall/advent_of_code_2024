package dailys

import (
	"aoc24/functions"
	"fmt"
	"regexp"
	"strconv"
)

func Day3_1(path string) int {
	fmt.Println("Running day 3 part 1 !")
	lines := functions.Readlines(path)

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	re2 := regexp.MustCompile(`[0-9]+`)
	sum := 0
	for _, line := range lines {
		matched := re.FindAllString(line, -1)
		//fmt.Printf("Found: %v\n", matched)
		for _, f := range matched {
			nums := re2.FindAllString(f, -1)
			v1, err1 := strconv.Atoi(nums[0])
			if err1 != nil {
				panic(err1)
			}
			v2, err2 := strconv.Atoi(nums[1])
			if err2 != nil {
				panic(err2)
			}
			sum += v1 * v2
		}
	}
	return sum
}

func Day3_2(path string) int {
	fmt.Println("Running day 3 part 2!")

	lines := functions.Readlines(path)

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	re2 := regexp.MustCompile(`[0-9]+`)
	sum := 0
	add := true
	for _, line := range lines {
		matched := re.FindAllString(line, -1)
		for _, pattern := range matched {
			if pattern == "do()" {
				add = true
			} else if pattern == "don't()" {
				add = false
			} else if add {
				nums := re2.FindAllString(pattern, -1)
				v1, err1 := strconv.Atoi(nums[0])
				if err1 != nil {
					panic(err1)
				}
				v2, err2 := strconv.Atoi(nums[1])
				if err2 != nil {
					panic(err2)
				}
				sum += v1 * v2
			}

		}
	}
	return sum
}
