package day4

import (
	"bufio"
	"sort"
	"strings"
)

func IsAnagram(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	chars1 := strings.Split(s1, "")
	chars2 := strings.Split(s2, "")

	sort.Strings(chars1)
	sort.Strings(chars2)

	if strings.Join(chars1, "") == strings.Join(chars2, "") {
		return true
	}

	return false
}

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

func isValid2(row string) bool {
	words := strings.Split(row, " ")
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if IsAnagram(words[i], words[j]) {
				return false
			}
		}
	}

	return true
}

//PartTwo returns num of valid passcodes
func PartTwo(input string) (result int) {

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		row := scanner.Text()
		if isValid2(row) {
			result++
		}
	}

	return
}
