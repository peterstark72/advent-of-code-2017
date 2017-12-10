package day3

import (
	"fmt"
	"math"
)

//Point is a coordinate point
type Point struct {
	X, Y int
}

//String returns printable version of Point
func (p Point) String() string {
	return fmt.Sprintf("X = %v, Y = %v", p.X, p.Y)
}

//Directions
const (
	Up int = iota
	Left
	Down
	Right
)

/*Spiral generates coordinates in a spiral pattern
from 0 to max

17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...
*/
func spiral(max int) chan Point {

	ch := make(chan Point)

	go func() {
		var p Point

		ch <- p

		steps := 1
		dir := Right
		size := 3
		for i := 1; i < max; i++ {
			steps--
			switch dir {
			case Right:
				p.X++
				if steps == 0 {
					dir = Up
					steps = size - 2
				}
			case Left:
				p.X--
				if steps == 0 {
					dir = Down
					steps = size - 1
				}
			case Up:
				p.Y++
				if steps == 0 {
					dir = Left
					steps = size - 1
				}
			case Down:
				p.Y--
				if steps == 0 {
					dir = Right
					steps = size
					size += 2
				}
			}
			ch <- p
		}
		close(ch)
	}()
	return ch
}

//PartOne gets the Manhattan distance from pos to origo
func PartOne(pos int) (dist int) {

	//Step through all points up to pos
	var point Point
	for p := range spiral(pos) {
		point = p
	}

	//Manhattan distance
	dist = int(math.Abs(float64(point.X)) + math.Abs(float64(point.Y)))

	return
}

//PartTwo returns the next sum larger than value
func PartTwo(value int, max int) (res int) {

	//Store all sum values in a map structure
	matris := make(map[int]map[int]int)

	//Init with 1 at (0,0)
	matris[0] = map[int]int{0: 1}
	points := spiral(max)
	//Skip first coordinate
	<-points

	for p := range points {
		res = 0
		if matris[p.X-1] != nil {
			res += matris[p.X-1][p.Y] + matris[p.X-1][p.Y-1] + matris[p.X-1][p.Y+1]
		}
		if matris[p.X+1] != nil {
			res += matris[p.X+1][p.Y] + matris[p.X+1][p.Y-1] + matris[p.X+1][p.Y+1]
		}
		if matris[p.X] != nil {
			res += matris[p.X][p.Y+1] + matris[p.X][p.Y-1]
		} else {
			matris[p.X] = make(map[int]int)
		}
		if res > value {
			//We have a result
			break
		}
		matris[p.X][p.Y] = res
	}

	return
}
