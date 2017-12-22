package day6

import "fmt"

//INPUT is the puzzle input
var INPUT = []int{2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14}

//Max returns index of the maximum value
func Max(values []int) (value int, index int) {

	for i, val := range values {
		if val > value {
			value = val
			index = i
		}
	}

	return
}

//redist redistributes the blocks
func redist(banks []int, history map[string]int) (cycles int, loops int) {

	cycles = len(history)

	//fmt.Println(cycles, banks)

	//Save configuration and check if we're done
	h := fmt.Sprintf("%v", banks)
	if history[h] > 0 {
		//This config has been seen before
		loops = cycles - history[h]
		return
	}
	history[h] = cycles

	//Get bank with most blocks and empty
	blocks, index := Max(banks)
	banks[index] = 0

	//Distribute blocks
	i := index
	length := len(banks)
	for {
		i = (i + 1) % length
		banks[i]++
		blocks--

		if blocks == 0 {
			break
		}
	}

	return redist(banks, history)
}

//PartOne
func PartOne(input []int) (counts int) {
	history := make(map[string]int)
	banks := make([]int, len(input))
	copy(banks, input)

	counts, _ = redist(banks, history)

	return
}

//PartTwo
func PartTwo(input []int) (loops int) {
	history := make(map[string]int)
	banks := make([]int, len(input))
	copy(banks, input)

	_, loops = redist(banks, history)

	return
}
