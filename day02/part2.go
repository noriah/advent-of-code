package day02

import "bytes"

func Part2(data []byte) int {
	posZ := 0
	posY := 0
	aim := 0

	splits := bytes.Split(data, []byte("\n"))

	for _, s := range splits {
		parts := bytes.Split(s, []byte(" "))
		switch parts[0][0] {
		case forward:
			v := int(parts[1][0] - '0')
			posY += v
			posZ += v * aim
		case up:
			aim -= int(parts[1][0] - '0')
		case down:
			aim += int(parts[1][0] - '0')
		default:
			panic(string(s))
		}
	}

	return posZ * posY
}
