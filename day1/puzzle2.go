package day1

import "fmt"

type Window struct {
	values []int
	sum    int
}

func NewWindow() *Window {
	return &Window{}
}

func (window *Window) AddElement(element int) bool {
	if len(window.values) > 2 {
		return false
	}

	window.values = append(window.values, element)
	window.sum += element

	return true
}

func (window *Window) CompareTo(otherWindow *Window) int {
	switch {
	case window.sum < otherWindow.sum:
		return -1
	case window.sum == otherWindow.sum:
		return 0
	default:
		return 1
	}
}

func (window *Window) String() string {
	return fmt.Sprintf("Window(%v (sum: %d))", window.values, window.sum)
}

func Puzzle2() int {
	var increments int
	var windows []*Window

	for _, value := range data {
		windows = append(windows, NewWindow())
		for _, window := range windows {
			window.AddElement(value)
		}

		if len(windows) > 3 {
			if windows[0].CompareTo(windows[1]) == -1 {
				increments++
			}
			windows = windows[1:]
		}
	}

	return increments
}
