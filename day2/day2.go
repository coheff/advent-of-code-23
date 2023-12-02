package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("day2/day2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	ans := 0

	// TODO: Parse out game number
	game := 1
	for fileScanner.Scan() {
		s := fileScanner.Text()
		colours := groupColours(s)

		if canPlay(colours) {
			ans += game
		}

		game++
	}

	fmt.Println(ans)
}

func groupColours(s string) map[string][]int {
	coloursToCounts := map[string][]int{"red": {}, "green": {}, "blue": {}}

	// remove spaces
	s = strings.ReplaceAll(s, " ", "")

	for colour := range coloursToCounts {
		coloursToCounts[colour] = getCountsForColour(colour, s)
	}

	return coloursToCounts
}

// uses strings.Index to find next index of colour occurance, parses number, trims string, & repeat
// easier approach would have been splitting on , or ;
func getCountsForColour(colour string, s string) []int {
	counts := []int{}

	for i := strings.Index(s, colour); i != -1 || len(s) == 0; i = strings.Index(s, colour) {
		// check for two digit count
		slice := s[i-2 : i]
		ansInt, err := strconv.Atoi(slice)
		if err == nil {
			// is a number
			counts = append(counts, ansInt)
		} else {
			// check for one digit count
			slice = s[i-1 : i]
			ansInt, err = strconv.Atoi(slice)
			if err == nil {
				// is a number
				counts = append(counts, ansInt)
			}
		}

		// trim input string
		s = s[i+1:]
	}

	return counts
}

func canPlay(colours map[string][]int) bool {
	bag := map[string]int{"red": 12, "green": 13, "blue": 14}

	for colour, counts := range colours {
		bagCount := bag[colour]

		for i := range counts {
			if counts[i] > bagCount {
				return false
			}
		}
	}

	return true
}
