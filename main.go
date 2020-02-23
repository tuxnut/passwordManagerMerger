package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Password struct {
	group    string // optional
	title    string // mandatory
	url      string // optional
	user     string // mandatory
	password string // mandatory
	notes    string // optional
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(file string) string {
	content, err := ioutil.ReadFile("./pc_pwd_db.txt")
	check(err)
	return string(content)
}

func parseDbG(data string) []Password {
	lines := strings.Split(data, "\n")

	nbOfPassword := len(lines) - 1
	passwordList := make([]Password, nbOfPassword)

	fmt.Println(lines[0])

	return passwordList
}

func main() {
	args := os.Args[2:] // first arg is prgm, second is '--' when using 'go run'

	if len(args) < 1 {
		panic(nil)
	}

	data := readFile(args[0])

	parseDbG(data)
}
