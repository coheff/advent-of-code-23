package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("day1/day1.txt")
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
