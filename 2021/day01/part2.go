package day01

import (
	"bytes"

	"github.com/noriah/advent-of-code/common"
)

func Part2(data []byte) int {
	window1 := &Window{
		Data:     make([]int, 3),
		Capacity: 3,
	}

	window2 := &Window{
		Data:     make([]int, 3),
		Capacity: 3,
	}

	splits := bytes.Split(data, []byte("\n"))

	// insert value 1
	window1.Push(common.ParseIntBytes(splits[0]))

	// insert values 2 and 3
	for _, s := range splits[1:3] {
		value := common.ParseIntBytes(s)
		window1.Push(value)
		window2.Push(value)
	}

	increases := 0

	// here we go
	for _, s := range splits[3:] {
		value := common.ParseIntBytes(s)
		window2.Push(value)

		if window2.Sum() > window1.Sum() {
			increases++
		}

		window1.Push(value)
	}

	return increases
}
