package day02

import "bytes"

func Part1(data []byte) int {
	posZ := 0
	posY := 0

	splits := bytes.Split(data, []byte("\n"))

	for _, s := range splits {
		parts := bytes.Split(s, []byte(" "))
		switch parts[0][0] {
		case forward:
			posY += int(parts[1][0] - '0')
		case up:
			posZ -= int(parts[1][0] - '0')
		case down:
			posZ += int(parts[1][0] - '0')
		default:
			panic(string(s))
		}
	}

	return posZ * posY
}
