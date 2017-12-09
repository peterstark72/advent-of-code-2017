package day2_test

import (
	"testing"

	"github.com/peterstark72/advent-of-code-2017/day2"
)

const input1 = `5	1	9	5
7	5	3
2	4	6	8`

const input2 = `5	9	2	8
9	4	7	3
3	8	6	5`

func TestOne(t *testing.T) {
	if day2.PartOne(input1) != 18 {
		t.Error("Should be 18")
	}
	if day2.PartOne(day2.INPUT) != 32020 {
		t.Error("Should be 32020")
	}
}

func TestTwo(t *testing.T) {
	if day2.PartTwo(input2) != 9 {
		t.Error("Should be 9")
	}
	if day2.PartTwo(day2.INPUT) != 236 {
		t.Error("Should be 236")
	}

}
