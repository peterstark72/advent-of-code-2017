package day4

import (
	"bufio"
	"sort"
	"strings"
)

//IsAnagram returns true of s1 and s2 are anagrams
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

//isvalid2 returns true if row is a valid passcode (part 2)
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
