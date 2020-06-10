package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const fileName = "phone-book.json"

type Person struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Age    uint   `json:"age"`
}

func readFromFile(filename string) (people []Person, err error) {
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

func writeToFile(filename string, peopleList []Person) error {
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

	fmt.Println(people)

	writeToFile(fileName, people)

	if err != nil {
		log.Fatal(err)
		return
	}
}
