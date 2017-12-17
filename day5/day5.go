package day5

import (
	"strconv"
	"strings"
)

func getnumbers(input string) (steps []int) {
	for _, s := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(s)
		steps = append(steps, num)
	}

	return
}

func PartOne(input string) (totalSteps int) {

	steps := getnumbers(input)
	length := len(steps)

	var current int
	var jump int
	for {
		totalSteps++
		jump = steps[current]

		steps[current]++

		current += jump

		if current >= length {
			break
		}

	}

	return
}

func PartTwo(input string) (totalSteps int) {

	steps := getnumbers(input)
	length := len(steps)

	var current int
	var jump int
	for {
		totalSteps++
		jump = steps[current]
		if steps[current] >= 3 {
			steps[current]--
		} else {
			steps[current]++
		}
		current += jump

		if current >= length {
			break
		}

	}

	return
}
