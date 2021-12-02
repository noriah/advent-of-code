package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/noriah/aoc2021/common"
	"github.com/noriah/aoc2021/day01"
)

var daySolutions = make([]common.DayRunner, 0, 25)

func init() {
	daySolutions = append(daySolutions, day01.NewRunner())
}

func main() {
	day := 0
	flag.IntVar(&day, "day", 1, "day for input")
	flag.Parse()

	if day < 1 || day > len(daySolutions) {
		panic("day out of range")
	}

	data := bytes.Trim(common.GetInput(day), "\n ")

	fmt.Printf("Day %d\n", day)

	dayRunner := daySolutions[day-1]
	dayRunner(data)
}
