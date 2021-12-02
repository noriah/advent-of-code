package day01

import (
	"bytes"

	"github.com/noriah/advent-of-code/common"
)

func Part1(data []byte) int {
	splits := bytes.Split(data, []byte("\n"))
	prev := common.ParseIntBytes(splits[0])
	increases := 0

	for _, s := range splits[1:] {
		value := common.ParseIntBytes(s)

		if prev < value {
			increases++
		}

		prev = value
	}

	return increases
}
