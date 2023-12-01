package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("day1two/day1two.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	ans := 0
	for fileScanner.Scan() {
		s := fileScanner.Text()
		ans = ans + calibrate(s)
	}
	fmt.Println(ans)
}

func calibrate(s string) int {
	ans := ""

	i := 0
	for {
		_, err := strconv.Atoi(string(s[i]))
		if err != nil {
			fmt.Println(string(s[i]) + " is not a number")

			//////////////////////////////////////////////
			// PART 2
			slice := s[0 : i+1]
			num := parseNumber(slice)
			if num != "" {
				fmt.Println(slice + " is a number")
				ans = num + ans
				break
			} else {
				fmt.Println(slice + " does not contain a number")
			}
			//////////////////////////////////////////////

			i++
		} else {
			fmt.Println(string(s[i]) + " is a number")
			ans = string(s[i]) + ans
			break
		}
	}

	j := len(s) - 1
	for {
		_, err := strconv.Atoi(string(s[j]))
		if err != nil {
			fmt.Println(string(s[j]) + " is not a number")

			//////////////////////////////////////////////
			// PART 2
			slice := s[j:]
			num := parseNumber(slice)
			if num != "" {
				fmt.Println(slice + " is a number")
				ans = ans + num
				break
			} else {
				fmt.Println(slice + " does not contain a number")
			}
			//////////////////////////////////////////////

			j--
		} else {
			fmt.Println(string(s[j]) + " is a number")
			ans = ans + string(s[j])
			break
		}
	}

	ansInt, err := strconv.Atoi(ans)
	if err != nil {
		fmt.Println("Uh-oh " + ans + " isn't a number...")
	}

	fmt.Println(ansInt)
	return ansInt
}

func parseNumber(s string) string {
	var lettersToNumber = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	for key, val := range lettersToNumber {
		if strings.Contains(s, key) {
			return val
		}
	}

	return ""
}
