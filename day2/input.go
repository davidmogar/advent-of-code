package day2

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

type command struct {
	Order string
	Units uint
}

var (
	//go:embed input
	input string

	Commands = func() (commands []command) {
		scanner := bufio.NewScanner(strings.NewReader(input))

		for scanner.Scan() {
			instruction := strings.Split(scanner.Text(), " ")
			if len(instruction) == 2 {
				units, err := strconv.Atoi(instruction[1])
				if err == nil {
					commands = append(commands, command{
						Order: instruction[0],
						Units: uint(units),
					})
				}
			}
		}

		return
	}()
)
