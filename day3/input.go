package day3

import (
	"bufio"
	_ "embed"
	"strings"
)

var (
	//go:embed input
	input string

	reports = func() (reports [][]int) {
		scanner := bufio.NewScanner(strings.NewReader(input))

		for scanner.Scan() {
			var report []int

			for _, value := range scanner.Text() {
				report = append(report, int(value)-48)
			}

			reports = append(reports, report)
		}

		return
	}()

	reportLength = func() int {
		if len(reports) > 0 {
			return len(reports[0])
		} else {
			return 0
		}
	}()
)
