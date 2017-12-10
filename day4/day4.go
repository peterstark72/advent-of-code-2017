package day4

import (
	"bufio"
	"strings"
)

//isValid returns true if row is a valid passcode
func isValid(row string) bool {

	seen := make(map[string]bool)
	words := strings.Split(row, " ")
	for _, w := range words {
		if seen[w] {
			return false
		}
		seen[w] = true
	}
	return true
}

//PartOne returns num of valid passcodes
func PartOne(input string) (result int) {

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		row := scanner.Text()
		if isValid(row) {
			result++
		}
	}

	return
}
