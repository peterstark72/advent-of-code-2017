package day1

//PartOne sum if value is same as next
func PartOne(s string) int {

	input := []rune(s)
	fch := int(input[0]) - '0'
	prev := fch
	var sum, d int
	var r rune
	for _, r = range input[1:] {
		d = int(r) - '0'
		if prev == d {
			sum += prev
		}
		prev = d
	}
	if fch == prev {
		sum += prev
	}
	return sum
}

//PartTwo sum if value is same as length/2 steps ahead
func PartTwo(s string) int {

	input := []rune(s)
	length := len(input)
	steps := length / 2
	var sum, d1, d2, j int
	for i, r := range input {
		d1 = int(r) - '0'
		j = (i + steps) % length
		d2 = int(input[j]) - '0'
		if d1 == d2 {
			sum += d1
		}
	}
	return sum
}
