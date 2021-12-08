package day6

import (
	_ "embed"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input string

	GetInitialState = func() *map[int]int {
		values := map[int]int{}

		for _, value := range strings.Split(input, ",") {
			timer, err := strconv.Atoi(value)
			if err == nil {
				if _, ok := values[timer]; ok {
					values[timer]++
				} else {
					values[timer] = 1
				}
			}
		}

		return &values
	}
)
