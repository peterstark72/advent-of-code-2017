package day5_test

import (
	"testing"

	"github.com/peterstark72/advent-of-code-2017/day5"
)

const Example = `0
3
0
1
-3`

func TestPartOne(t *testing.T) {

	if day5.PartOne(Example) != 5 {
		t.Error("Part One")
	}

}

func TestPartTwo(t *testing.T) {

	if day5.PartTwo(Example) != 10 {
		t.Error("Part Two")
	}

}
