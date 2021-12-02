package common

import "strconv"

func ParseIntBytes(value []byte) int {
	i, err := strconv.ParseInt(string(value), 10, 64)
	CheckError(err)

	return int(i)
}
