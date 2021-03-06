package main

import (
	"fmt"

	"github.com/peterstark72/advent-of-code-2017/day1"
	"github.com/peterstark72/advent-of-code-2017/day2"
	"github.com/peterstark72/advent-of-code-2017/day3"
	"github.com/peterstark72/advent-of-code-2017/day4"
	"github.com/peterstark72/advent-of-code-2017/day5"
	"github.com/peterstark72/advent-of-code-2017/day6"
)

func main() {

	fmt.Println("Day 1 / Part 1: ", day1.PartOne(day1.INPUT))
	fmt.Println("Day 1 / Part 2: ", day1.PartTwo(day1.INPUT))

	fmt.Println("Day 2 / Part 1: ", day2.PartOne(day2.INPUT))
	fmt.Println("Day 2 / Part 2: ", day2.PartTwo(day2.INPUT))

	fmt.Println("Day 3 / Part 1: ", day3.PartOne(347991))
	fmt.Println("Day 3 / Part 2: ", day3.PartTwo(347991, 1024))

	fmt.Println("Day 4 / Part 1: ", day4.PartOne(day4.INPUT))
	fmt.Println("Day 4 / Part 2: ", day4.PartTwo(day4.INPUT))

	fmt.Println("Day 5 / Part 1: ", day5.PartOne(day5.INPUT))
	fmt.Println("Day 5 / Part 2: ", day5.PartTwo(day5.INPUT))

	fmt.Println("Day 6 / Part 1: ", day6.PartOne(day6.INPUT))
	fmt.Println("Day 6 / Part 2: ", day6.PartTwo(day6.INPUT))

}
