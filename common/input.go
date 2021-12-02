package common

import (
	"fmt"
	"os"
)

const pathTmpl = "%d/input/day%02d.txt"

func GetInput(year, day int) []byte {
	filePath := fmt.Sprintf(pathTmpl, year, day)

	if _, err := os.Stat(filePath); err == nil {
		data, err := os.ReadFile(filePath)
		CheckError(err)
		return data
	}

	data := FetchInput(year, day)

	err := os.WriteFile(filePath, data, 0644)
	CheckError(err)

	return data
}
