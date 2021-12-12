package day3

import (
	"math"
)

func Puzzle1() int64 {
	count := make([]int, reportLength)

	for _, report := range reports {
		for i, value := range report {
			count[i] += value
		}
	}

	target := len(reports) / 2
	var epsilonRate, gammaRate float64

	for i, value := range count {
		if value > target { // more 1s
			gammaRate += math.Pow(2, float64(reportLength - i - 1))
		} else { // more 0s
			epsilonRate += math.Pow(2, float64(reportLength - i - 1))
		}
	}

	return int64(epsilonRate * gammaRate)
}
