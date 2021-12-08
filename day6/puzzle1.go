package day6

import (
	_ "embed"
)

func Puzzle1() int {
	return simulate(GetInitialState(), 79)
}
