package day6_test

import (
	"testing"

	"github.com/peterstark72/advent-of-code-2017/day6"
)

func TestMax(t *testing.T) {
	v, i := day6.Max([]int{1, 2, 3, 4, 4})
	if (v != 4) || (i != 3) {
		t.Error("Max")
	}
	v, i = day6.Max([]int{7, 3, 9, 11, 4})
	if (v != 11) || (i != 3) {
		t.Error("Max")
	}
	v, i = day6.Max([]int{12, 12, 9, 11, 4})
	if (v != 12) || (i != 0) {
		t.Error("Max")
	}
}

func TestOne(t *testing.T) {
	banks := []int{0, 2, 7, 0}
	if day6.PartOne(banks) != 5 {
		t.Error("Part one")
	}
}

func TestTwo(t *testing.T) {
	banks := []int{0, 2, 7, 0}
	if day6.PartTwo(banks) != 4 {
		t.Error("Part Two")
	}
}
