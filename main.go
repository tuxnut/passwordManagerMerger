package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	TextGroupField    = 1
	TextTitleField    = 2
	TextURLField      = 3
	TextUserField     = 4
	TextPasswordField = 5
	TextNotesField    = 6
)

const (
	CSVTitleField    = 0
	CSVCategoryField = 1
	CSVUserField     = 2
	CSVPasswordField = 3
	CSVURLField      = 4
	CSVCommentsField = 5
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

func readTextFile(file string) []string {
	content, err := ioutil.ReadFile(file)
	check(err)
	return strings.Split(string(content), "\n")
}

func parseDbG(dataLines []string) []Password {
	// Remove last line if empty
	if "" == dataLines[len(dataLines)-1] {
		dataLines = dataLines[:len(dataLines)-1]
	}

	nbOfPassword := len(dataLines) - 1
	passwordList := make([]Password, nbOfPassword)

	for i := 0; i < nbOfPassword; i++ {
		passwordFields := strings.Split(dataLines[i+1], ",")

		if len(passwordFields) >= TextNotesField {
			passwordList[i] = Password{
				group:    passwordFields[TextGroupField],
				title:    passwordFields[TextTitleField],
				url:      passwordFields[TextURLField],
				user:     passwordFields[TextUserField],
				password: passwordFields[TextPasswordField],
				notes:    passwordFields[TextNotesField],
			}
		}
	}

	return passwordList
}

func parseCSVFile(file string) []Password {
	csvFile, _ := os.Open(file)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var passwordList []Password
	for {
		passwordFields, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		passwordList = append(passwordList, Password{
			title:    passwordFields[CSVTitleField],
			group:    passwordFields[CSVCategoryField],
			user:     passwordFields[CSVUserField],
			password: passwordFields[CSVPasswordField],
			url:      passwordFields[CSVURLField],
			notes:    passwordFields[CSVCommentsField],
		})
	}

	return passwordList
}

func compare(left, right Password) bool {
	return left.title == right.title &&
		left.user == right.user &&
		left.password == right.password
}

func retrieveDuplicates(left, right []Password) []Password {
	var duplicates []Password
	for i := 0; i < len(left); i++ {
		currPassword := left[i]

		for j := 0; j < len(right); j++ {
			if compare(currPassword, right[j]) {
				duplicates = append(duplicates, currPassword)
			}
		}
	}
	return duplicates
}

func main() {
	args := os.Args[2:] // first arg is prgm, second is '--' when using 'go run'

	if len(args) < 1 {
		panic(nil)
	}

	var passwordList [][]Password

	for i := 0; i < len(args); i++ {
		if strings.HasSuffix(args[i], "txt") {
			dataLines := readTextFile(args[i])
			passwordList = append(passwordList, parseDbG(dataLines))

			// for j := 0; j < len(passwordList[i]); j++ {
			// 	fmt.Println(passwordList[i][j].title)
			// }
		} else if strings.HasSuffix(args[i], "csv") {
			passwordList = append(passwordList, parseCSVFile(args[i]))

			// for j := 0; j < len(passwordList[i]); j++ {
			// 	fmt.Println(passwordList[i][j].title)
			// }
		}
	}

	duplicates := retrieveDuplicates(passwordList[0], passwordList[1])

	fmt.Println(len(duplicates))
	for i := 0; i < len(duplicates); i++ {
		fmt.Println(duplicates[i].title)
	}
}
