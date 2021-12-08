package day6

import (
	_ "embed"
)

func Puzzle2() int {
	return simulate(GetInitialState(), 255)
}
