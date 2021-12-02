package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/noriah/advent-of-code/common"

	"github.com/noriah/advent-of-code/2020/day17"
)

var daySolutions = make([]common.DayRunner, 25)

func init() {
	daySolutions[16] = day17.NewRunner()
}

func main() {
	day := 0
	flag.IntVar(&day, "day", 1, "day for input")
	flag.Parse()

	if day < 1 || day > len(daySolutions) {
		panic("day out of range")
	}

	data := bytes.Trim(common.GetInput(2020, day), "\n ")

	fmt.Printf("Day %d\n", day)

	dayRunner := daySolutions[day-1]
	dayRunner(data)
}
