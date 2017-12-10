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

}
