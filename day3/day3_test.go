package day3_test

import (
	"testing"

	"github.com/peterstark72/advent-of-code-2017/day3"
)

func TestOne(t *testing.T) {
	if day3.PartOne(23) != 2 {
		t.Error("23 should be 2")
	}
	if day3.PartOne(12) != 3 {
		t.Error("12 should be 3")
	}
	if day3.PartOne(13) != 4 {
		t.Error("13 should be 4")
	}
	if day3.PartOne(1024) != 31 {
		t.Error("1024 should be 31")
	}
	if day3.PartOne(14) != 3 {
		t.Error("14 should be 3")
	}
	if day3.PartOne(10) != 3 {
		t.Error("10 should be 3")
	}
	if day3.PartOne(25) != 4 {
		t.Error("25 should be 4")
	}
	if day3.PartOne(1) != 0 {
		t.Error("1 should be 0")
	}
}

func TestTwo(t *testing.T) {

	const MAX = 1024

	if day3.PartTwo(1, MAX) != 2 {
		t.Error("1 should be 2")
	}
	if day3.PartTwo(4, MAX) != 5 {
		t.Error("4 should be 5")
	}
	if day3.PartTwo(351, MAX) != 362 {
		t.Error("351 should be 3625")
	}
	if day3.PartTwo(347991, MAX) != 349975 {
		t.Error("347991 should be 349975")
	}

}
