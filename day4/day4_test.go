package day4_test

import (
	"testing"

	"github.com/peterstark72/advent-of-code-2017/day4"
)

func TestOne(t *testing.T) {

	if day4.PartOne("aa bb cc dd ee") != 1 {
		t.Error("aa bb cc dd ee")
	}
	if day4.PartOne("aa bb cc dd aa") != 0 {
		t.Error("aa bb cc dd aa")
	}
	if day4.PartOne("aa bb cc dd aaa") != 1 {
		t.Error("aa bb cc dd aaa")
	}
	if day4.PartOne(day4.INPUT) != 477 {
		t.Error("INPUT")
	}
}

func TestIsAnagram(t *testing.T) {
	if day4.IsAnagram("abdc", "dcab") != true {
		t.Error("abcd")
	}
	if day4.IsAnagram("abdw", "dcabm") != false {
		t.Error("abdw")
	}
}

func TestTwo(t *testing.T) {

	if day4.PartTwo("abcde fghij") != 1 {
		t.Error("abcde fghij")
	}
	if day4.PartTwo("abcde xyz ecdab") != 0 {
		t.Error("abcde xyz ecdab")
	}
	if day4.PartTwo("a ab abc abd abf abj") != 1 {
		t.Error("a ab abc abd abf abj")
	}
	if day4.PartTwo("iiii oiii ooii oooi oooo") != 1 {
		t.Error("iiii oiii ooii oooi oooo")
	}
	if day4.PartTwo("oiii ioii iioi iiio") != 0 {
		t.Error("oiii ioii iioi iiio")
	}
}
