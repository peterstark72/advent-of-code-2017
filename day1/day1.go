package day1

//PartOne sum if value is same as next
func PartOne(input string) (sum int) {

	numbers := []rune(input)
	length := len(numbers)

	for i, n := range numbers {

		if n == numbers[(i+1)%length] {
			sum += int(n - '0')
		}

	}

	return
}

//PartTwo sum if value is same as length/2 steps ahead
func PartTwo(input string) (sum int) {

	numbers := []rune(input)
	length := len(numbers)
	steps := length / 2

	for i, n := range numbers {

		if n == numbers[(i+steps)%length] {
			sum += int(n - '0')
		}

	}

	return
}
