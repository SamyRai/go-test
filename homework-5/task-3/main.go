package main

import (
	"github.com/SamyRai/go-test/homework-5/task-3/csver"
)

// Info about applicants for photo shooting
type Applicant struct {
	Name   string `csv:"name"`
	Phone  string `csv:"phone"`
	Age    uint   `csv:"age"`
	Gender string `csv:"gender"`
	Height uint   `csv:"height"`
}

// Container for multiple applicants
type Applicants []Applicant

func (applicants *Applicants) fillFakeData() {
	newData := Applicants{
		Applicant{
			Name:   "Viktor",
			Phone:  "234234234",
			Age:    23,
			Gender: "male",
			Height: 177,
		},
		Applicant{
			Name:   "Hanna",
			Phone:  "9494949",
			Age:    18,
			Gender: "female",
			Height: 183,
		},
		Applicant{
			Name:   "Sara",
			Phone:  "2929934",
			Age:    30,
			Gender: "female",
			Height: 168,
		},
	}

	*applicants = append(*applicants, newData...)
}

func main() {
	applicants := *new(Applicants)
	applicants.fillFakeData()
	applicantsAsInterface := make([]interface{}, len(applicants))
	for i, v := range applicants {
		applicantsAsInterface[i] = v
	}
	csver.WriteToFile(applicantsAsInterface, "test.csv")
}
