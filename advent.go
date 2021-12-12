package main

import (
	"advent/day1"
	day2puzzle1 "advent/day2/puzzle1"
	day2puzzle2 "advent/day2/puzzle2"
	"advent/day3"
	"advent/day6"
	"fmt"
)

func main() {
	fmt.Println("Day 1:")
	fmt.Println("  - Puzzle 1:", day1.Puzzle1())
	fmt.Println("  - Puzzle 2:", day1.Puzzle2())
	fmt.Println("Day 2:")
	fmt.Println("  - Puzzle 1:", day2puzzle1.Puzzle1())
	fmt.Println("  - Puzzle 2:", day2puzzle2.Puzzle2())
	fmt.Println("Day 3:")
	fmt.Println("  - Puzzle 1:", day3.Puzzle1())
	fmt.Println("  - Puzzle 2:", day3.Puzzle2())
	fmt.Println("Day 6:")
	fmt.Println("  - Puzzle 1:", day6.Puzzle1())
	fmt.Println("  - Puzzle 2:", day6.Puzzle2())
}
