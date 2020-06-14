package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

const fileName = "phone-book.json"

type Person struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Age    uint   `json:"age"`
}

type PhoneBook []Person

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (p PhoneBook) Len() int {
	return len(p)
}

func (p PhoneBook) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p PhoneBook) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func readFromFile(filename string) (people PhoneBook, err error) {
	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		return people, err
	}

	err = json.Unmarshal(jsonString, &people)
	if err != nil {
		return people, err
	}

	return people, err
}

func writeToFile(filename string, peopleList PhoneBook) error {
	peopleInJson, err := json.Marshal(peopleList)
	if err != nil {
		return err
	}

	peopleInByte := []byte(peopleInJson)
	err = ioutil.WriteFile(filename, peopleInByte, 0755)

	return nil
}

func main() {
	people, err := readFromFile(fileName)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(people)

	people = append(people, Person{"New Person", "1250232565", 33})
	people = append(people, Person{"Ali Baba", "1250232565", 33})

	fmt.Println("Before sorting", people)
	sort.Sort(people)
	fmt.Println("After sorting", people)

	writeToFile(fileName, people)

	if err != nil {
		log.Fatal(err)
		return
	}
}
