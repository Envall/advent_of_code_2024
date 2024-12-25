package dailys

import (
	"aoc24/functions"
	"fmt"
)

type coord struct {
	x int
	y int
}
type dir struct {
	dx int
	dy int
}

var dirsMap = map[string]dir{
	"N":  {dx: 0, dy: -1},  // N
	"NE": {dx: 1, dy: -1},  // NE
	"E":  {dx: 1, dy: 0},   // E
	"SE": {dx: 1, dy: 1},   // SE
	"S":  {dx: 0, dy: 1},   // S
	"SW": {dx: -1, dy: 1},  // SW
	"W":  {dx: -1, dy: 0},  // W
	"NW": {dx: -1, dy: -1}, // NW
}

//dirsMap["N"] = {dx: 0, dy: -1},  // N

func findWord(r int, c int, word string, lines []string) (bool, []string) {
	// Starting from lines[r][c], see if 'MAS' is spelled
	// in any direction

	rmax := len(lines) - 1
	cmax := len(lines[0]) - 1
	wlen := len(word)
	//fmt.Printf("r:%d c:%d\n", r, c)
	//fmt.Printf("rmax:%d cmax:%d\n", rmax, cmax)

	var okdirs []string
	for dirname, dir := range dirsMap {
		//fmt.Printf("Dir: %s\n", dirname)
		dx := dir.dx
		dy := dir.dy
		// Too close to first/last row?
		if r+(wlen*dy) > rmax || r+(wlen*dy) < 0 {
			continue
		}
		// Too close to first/last column?
		if c+(wlen*dx) > cmax || c+(wlen*dx) < 0 {
			continue
		}
		// Do letters come in the right order?
		wordfound := true
		for i, letter := range word {
			//found := lines[r+((i+1)*dy)][c+((i+1)*dx)]
			//fmt.Printf("letter:%v found:%v\n", letter, found)
			if lines[r+((i+1)*dy)][c+((i+1)*dx)] != byte(letter) {
				//fmt.Printf("Wrong letter found!\n")
				wordfound = false
				break
			}
		}
		if wordfound {
			okdirs = append(okdirs, dirname)
		}
	}

	if len(okdirs) > 0 {
		return true, okdirs
	} else {
		return false, okdirs
	}
}

func Day4_1(path string) int {
	/*
	   MMMSXXMASM
	   MSAMXMSMSA
	   AMXSXMAAMM
	   MSAMASMSMX
	   XMASAMXAMM
	   XXAMMXXAMA
	   SMSMSASXSS
	   SAXAMASAAA
	   MAMMMXMMMM
	   MXMXAXMASX

	   Strategy:
	   When X is found, look in all directions for M, if M is found then
	   continue in the same direction for A, S.
	*/

	fmt.Println("Running day 4 part 1 !")
	lines := functions.Readlines(path)
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println("----------")

	nxmas := 0
	for r := range len(lines) {
		//fmt.Println(lines[r])
		for c := range len(lines[0]) {
			//fmt.Printf("%c,", lines[r][c])

			letter := lines[r][c]
			if letter == 'X' {
				//fmt.Printf("Found X at row:%d col:%d\n", r, c)
				found, dirs := findWord(r, c, "MAS", lines)
				if found {
					//fmt.Printf("Found XMAS at row %d col %d, directions %v\n", r, c, dirs)
					nxmas += len(dirs)
				}
			}
		}
	}

	return nxmas
}

func findWordDiagonal(r int, c int, word string, lines []string) (bool, []string) {
	_, dirs := findWord(r, c, word, lines)

	diagdirs := []string{}
	if functions.StringInSlice(dirs, "NE") {
		diagdirs = append(diagdirs, "NE")
	}
	if functions.StringInSlice(dirs, "SE") {
		diagdirs = append(diagdirs, "SE")
	}
	if functions.StringInSlice(dirs, "SW") {
		diagdirs = append(diagdirs, "SW")
	}
	if functions.StringInSlice(dirs, "NW") {
		diagdirs = append(diagdirs, "NW")
	}

	found := false
	if len(diagdirs) > 0 {
		found = true
	}
	return found, diagdirs
}

func coordInSlice(c coord, s []coord) bool {
	for _, cs := range s {
		if c.x == cs.x && c.y == cs.y {
			return true
		}
	}
	return false
}

func Day4_2(path string) int {
	/* In this part we want to find "MAS" spelled as a cross
	for example:  M.S
	              .A.
				  M.S

	All directions are ok.

	Strategy:
	- iterate over all letters
	- if 'M' is found, try to find 'AS' in a diagonal direction
	- save location of 'A'
	- count number of 'A' locations that occur twice
	*/

	fmt.Println("Running day 4 part 2 !")
	lines := functions.Readlines(path)

	//for _, line := range lines {
	//	fmt.Println(line)
	//}
	//fmt.Println("----------")

	savedCoords := []coord{}
	nfound := 0
	for r := range len(lines) {
		for c := range len(lines[0]) {
			letter := lines[r][c]
			if letter == 'M' {
				found, dirs := findWordDiagonal(r, c, "AS", lines)
				if found {
					//fmt.Printf("Found 'AS' at row:%d col:%d dirs:%v\n", r, c, dirs)
					for _, dirname := range dirs {
						dir := dirsMap[dirname]
						rA := r + dir.dy
						cA := c + dir.dx
						coordA := coord{x: cA, y: rA}
						//fmt.Printf("Found 'A' at row:%d col:%d (middle of 'MAS')\n", rA, cA)
						if coordInSlice(coordA, savedCoords) {
							//fmt.Println("This coord was already in the list!")
							nfound += 1
						}
						savedCoords = append(savedCoords, coordA)
					}
				}
			}
		}
	}

	return nfound
}
