package day1

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input string

	data = func() (values []int) {
		scanner := bufio.NewScanner(strings.NewReader(input))

		for scanner.Scan() {
			value, err := strconv.Atoi(scanner.Text())
			if err == nil {
				values = append(values, value)
			}
		}

		return
	}()
)
