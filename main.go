package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	GroupField    = 1
	TitleField    = 2
	URLField      = 3
	UserField     = 4
	PasswordField = 5
	NotesField    = 6
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

	// Remove last line if empty
	if "" == lines[len(lines) - 1] {
		lines = lines[:len(lines) - 1]
	}

	nbOfPassword := len(lines) - 1
	passwordList := make([]Password, nbOfPassword)

	for i := 0; i < nbOfPassword; i++ {
		passwordFields := strings.Split(lines[i+1], ",")

		if len(passwordFields) >= 6 {
			passwordList[i] = Password{
				group:    passwordFields[GroupField],
				title:    passwordFields[TitleField],
				url:      passwordFields[URLField],
				user:     passwordFields[UserField],
				password: passwordFields[PasswordField],
				notes:    passwordFields[NotesField],
			}
		}
	}

	return passwordList
}

func main() {
	args := os.Args[2:] // first arg is prgm, second is '--' when using 'go run'

	if len(args) < 1 {
		panic(nil)
	}

	data := readFile(args[0])

	fmt.Println(parseDbG(data))
}
