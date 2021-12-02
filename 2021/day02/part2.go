package day02

import "bytes"

func Part2(data []byte) int {
	posZ := 0
	posY := 0
	aim := 0

	splits := bytes.Split(data, []byte("\n"))

	for _, s := range splits {
		parts := bytes.Split(s, []byte(" "))
		value := int(parts[1][0] - '0')

		switch parts[0][0] {
		case forward:
			posY += value
			posZ += value * aim
		case up:
			aim -= value
		case down:
			aim += value
		default:
			panic(string(s))
		}
	}

	return posZ * posY
}
