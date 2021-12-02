package day02

import "fmt"

const forward = 'f'
const up = 'u'
const down = 'd'

func NewRunner() func(data []byte) {
	return func(data []byte) {
		fmt.Printf("Problem 1: %d\n", Part1(data))
		fmt.Printf("Problem 2: %d\n", Part2(data))
	}
}
