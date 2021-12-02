package day01

import (
	"bytes"

	"github.com/noriah/advent-of-code/common"
)

func Part2(data []byte) int {
	window := &Window{
		Data:     make([]int, 3),
		Capacity: 3,
	}

	splits := bytes.Split(data, []byte("\n"))

	// insert values 1, 2 and 3
	for _, s := range splits[0:3] {
		window.Push(common.ParseIntBytes(s))
	}

	increases := 0

	// here we go
	for _, s := range splits[3:] {

		if window.Sum() < window.Push(common.ParseIntBytes(s)) {
			increases++
		}
	}

	return increases
}
