package day02

import "bytes"

func Part1(data []byte) int {
	posZ := 0
	posY := 0

	splits := bytes.Split(data, []byte("\n"))

	for _, s := range splits {
		parts := bytes.Split(s, []byte(" "))
		value := int(parts[1][0] - '0')

		switch parts[0][0] {
		case forward:
			posY += value
		case up:
			posZ -= value
		case down:
			posZ += value
		default:
			panic(string(s))
		}
	}

	return posZ * posY
}
