package functions

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Readlines(path string) []string {
	fmt.Printf("Reading line by line as strings: %s\n", path)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func Readf_2cols_int(path string) ([]int, []int) {
	fmt.Printf("Reading 2col int file: %s\n", path)

	// Read the contents of the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Closes when Readf_2cols_int is finished !

	scanner := bufio.NewScanner(file)

	// Create arrays to store results
	c1 := []int{}
	c2 := []int{}

	// Iterate through each line of file
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields((line))

		if len(words) > 2 {
			panic(fmt.Sprintf("Bad input data! More than 2 columns: %v", words))
		}

		v1, err1 := strconv.Atoi(words[0])
		if err1 != nil {
			panic(err1)
		}
		v2, err2 := strconv.Atoi(words[1])
		if err2 != nil {
			panic(err2)
		}

		c1 = append(c1, v1)
		c2 = append(c2, v2)
	}

	// TODO debug print n first characters of each array
	n1 := int(math.Min(float64(len(c1)), 10))
	fmt.Printf("c1: %d + [...]\n", c1[:n1])
	n2 := int(math.Min(float64(len(c2)), 10))
	fmt.Printf("c2: %v + [...]\n", c2[:n2])
	fmt.Printf("\n")
	return c1, c2
}

// Read file where each line is a string af whitespace-separated integers.
// Function returns a 2D slice of integers.
func Readf_lines_as_int_slices(path string) [][]int {
	fmt.Println(path)
	fmt.Printf("Reading input data into 2D slice of ints\n")

	// Read the contents of the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Closes when Readf_2cols_int is finished !

	var data [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Fields(line)
		numbers := []int{}
		for _, char := range chars {
			num, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
		}
		data = append(data, numbers)
	}

	return data
}

/*
func main() {
	file := "../input_data/day1_test.txt"
	c1, c2 := Readf_2cols_int((file))

	fmt.Printf("Column 1: %v\n", c1)
	fmt.Printf("Column 2: %v\n", c2)

	fmt.Println("goodbye")
}
*/
