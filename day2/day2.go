package day2

import (
	"math"
	"strconv"
	"strings"
)

func findEvenDivisibles(line []string) int {

	length := len(line)
	for i, a := range line[:length-1] {
		d1, _ := strconv.Atoi(a)
		for _, b := range line[i+1:] {
			d2, _ := strconv.Atoi(b)
			if (d1 % d2) == 0 {
				return d1 / d2
			}
			if (d2 % d1) == 0 {
				return d2 / d1
			}
		}
	}

	return 0
}

//PartTwo sum even divisable numbers for each row
func PartTwo(input string) (sum int) {

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sum += findEvenDivisibles(strings.Split(line, "\t"))
	}

	return
}

//PartOne sums diff for each row
func PartOne(input string) (sum int) {

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		numbers := strings.Split(line, "\t")

		max := 0
		min := math.MaxInt64
		for _, num := range numbers {
			d, _ := strconv.Atoi(num)
			if d > max {
				max = d
			}
			if d < min {
				min = d
			}
		}
		sum += max - min
	}

	return
}
