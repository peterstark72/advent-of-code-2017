package day1_test

import (
	"testing"

	"github.com/peterstark72/aoc2017/day1"
)

func TestPartOne(t *testing.T) {
	if day1.PartOne("1122") != 3 {
		t.Error("1122 should be 3")
	}
	if day1.PartOne("1111") != 4 {
		t.Error("1111 should be 4")
	}
	if day1.PartOne("1234") != 0 {
		t.Error("1234 should be 0")
	}
	if day1.PartOne("91212129") != 9 {
		t.Error("91212129 should be 9")
	}
	if day1.PartOne(day1.INPUT) != 1253 {
		t.Error("Should be 1253")
	}
}

func TestPartTwo(t *testing.T) {
	if day1.PartTwo("1212") != 6 {
		t.Error("1212 should be 6")
	}
	if day1.PartTwo("1221") != 0 {
		t.Error("1221 should be 0")
	}
	if day1.PartTwo("123425") != 4 {
		t.Error("123425 should be 4")
	}
	if day1.PartTwo("123123") != 12 {
		t.Error("123123 should be 12")
	}
	if day1.PartTwo("12131415") != 4 {
		t.Error("12131415 should be 4")
	}
	if day1.PartTwo(day1.INPUT) != 1278 {
		t.Error("Should be 1278")
	}
}
