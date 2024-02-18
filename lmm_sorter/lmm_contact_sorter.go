package lmm

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
	email     string
	wholeLine string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var people []Person

func main() {
	file, err := os.Open("lmm_emails.txt")
	check(err)
	defer file.Close()

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)
	var lines []string

	for reader.Scan() {
		lines = append(lines, reader.Text())
	}
	file.Close()
	for _, line := range lines {
		name := strings.Split(line, "|")[0]
		email := strings.Split(line, "|")[1]
		people = append(people, Person{email: email, firstName: strings.Split(name, " ")[0], lastName: strings.Split(name, " ")[1]})
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].lastName < people[j].lastName
	})
	// fmt.Println(people)
	for _, person := range people {
		fmt.Printf("\"%s %s|%s\" ", person.firstName, person.lastName, person.email)
	}

}
