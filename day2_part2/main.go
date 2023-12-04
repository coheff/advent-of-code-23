package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("day2_part2/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	ans := 0

	for fileScanner.Scan() {
		s := fileScanner.Text()
		colours := groupColours(s)

		//////////////////////////////////////////////
		// PART 2
		power := 1

		for _, counts := range colours {
			if len(counts) > 0 {
				power *= getMaxCount(counts)
			}
		}

		ans += power
		//////////////////////////////////////////////
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

func getMaxCount(counts []int) int {
	max := 1

	for i := range counts {
		if counts[i] > max {
			max = counts[i]
		}
	}

	return max
}
