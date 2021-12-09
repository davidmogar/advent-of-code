package puzzle2

import "advent/day2"

type submarine struct {
	aim                uint
	depth              uint
	horizontalPosition uint
}

func (sub *submarine) down(units uint) {
	sub.aim += units
}

func (sub *submarine) forward(units uint) {
	sub.depth += sub.aim * units
	sub.horizontalPosition += units
}

func (sub *submarine) up(units uint) {
	sub.aim -= units
}

func Puzzle2() uint {
	sub := submarine{}

	for _, command := range day2.Commands {
		switch command.Order {
		case "down":
			sub.down(command.Units)
		case "forward":
			sub.forward(command.Units)
		case "up":
			sub.up(command.Units)
		}
	}

	return sub.depth * sub.horizontalPosition
}
