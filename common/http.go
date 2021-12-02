package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const cookiePath = "session.cookie"

func readCookie() string {
	file, err := os.Open(cookiePath)
	CheckError(err)
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	CheckError(err)

	return strings.Trim(string(data), "\n ")
}

const url = "https://adventofcode.com/2021/day/%d/input"

func GetInput(day int) []byte {
	req, err := http.NewRequest("GET", fmt.Sprintf(url, day), nil)
	CheckError(err)

	sessionCookie := readCookie()

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := http.Client{
		Timeout: time.Duration(3) * time.Second,
	}

	resp, err := client.Do(req)
	CheckError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	return body
}
