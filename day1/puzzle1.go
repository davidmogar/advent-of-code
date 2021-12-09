package day1

import "math"

func Puzzle1() int {
	var increments int
	previous := math.MaxInt64

	for _, value := range data {
		if value > previous {
			increments++
		}
		previous = value
	}

	return increments
}
