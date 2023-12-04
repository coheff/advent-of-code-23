package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	number  string
	indices [][]int
}

// main function
func main() {
	ans := 0
	grid := readPuzzleInput("day3/input.txt")
	numbers := []Number{}
	// iterate over 2d array and populate list of {number, [index of each digit]}
	// e.g. {"467", [[0,0],[0,1],[0,2]]}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if isValidNumber(grid, row, col) {
				numBuilder := string(grid[row][col])
				indices := [][]int{{row, col}}
				nextCol := col + 1
				// while next rune is a number add to numBuilder and store index
				// TODO: duplication w/ previous lines - could be recursion
				for isValidNumber(grid, row, nextCol) {
					numBuilder += string(grid[row][nextCol])
					indices = append(indices, []int{row, nextCol})
					nextCol++
				}
				// add found number and indices to numbers list
				numbers = append(numbers, Number{numBuilder, indices})
				// move to end of found number
				col = nextCol
			}
		}
	}
	fmt.Println(numbers)
	// iterate over numberByIndices and check if symbol adjacent to indices
	for _, number := range numbers {
		for i := range number.indices {
			if hasSymbolAdjacent(grid, number.indices[i]) {
				// fmt.Printf("%v has a symbol adjacent", number.number)
				fmt.Println()
				n, err := strconv.Atoi(number.number)
				if err != nil {
					fmt.Println("Oops, " + number.number + " is not a number")
				}
				ans += n
				break
			}
		}
	}

	fmt.Println(ans)
}

// readPuzzleInput handles reading input file
// returns a 2d rune array
func readPuzzleInput(input string) [][]rune {
	readFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	var result [][]rune
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		row := []rune(line)
		result = append(result, row)
	}
	return result
}

// isValidNumber returns true if given coordinates for a grid are a number
// else returns false
func isValidNumber(grid [][]rune, row int, col int) bool {
	if row >= len(grid) || col >= len(grid[row]) {
		return false
	}
	return unicode.IsDigit(grid[row][col])
}

// hasSymbolAdjacent returns true if there's a symbol adjacent, including diagonally
// else returns false
func hasSymbolAdjacent(grid [][]rune, indices []int) bool {
	return isSymbol(grid, indices[0]-1, indices[1]-1) ||
		isSymbol(grid, indices[0]-1, indices[1]) ||
		isSymbol(grid, indices[0]-1, indices[1]+1) ||
		isSymbol(grid, indices[0], indices[1]-1) ||
		isSymbol(grid, indices[0], indices[1]+1) ||
		isSymbol(grid, indices[0]+1, indices[1]-1) ||
		isSymbol(grid, indices[0]+1, indices[1]) ||
		isSymbol(grid, indices[0]+1, indices[1]+1)
}

// isSymbol returns true if given coordinates for a grid are a symbol e.g. `*`, `$`, etc.
// else returns false
func isSymbol(grid [][]rune, row int, col int) bool {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) {
		return false
	}
	c := grid[row][col]
	return !unicode.IsDigit(c) && c != '.'
}
