package handler

import (
	"fmt"

	"github.com/refqihussein/golangforlife/golangforlife/Introduction/model"
)

//Create adalah
func Create() *model.Person {

	person := &model.Person{}
	dream := &model.Dream{}

	dream = &model.Dream{
		Type:        "basah",
		Description: "enak",
	}

	person.ID = 1
	person.Age = 23
	person.Name = "Refqi Hussein"
	person.Address = "cakung"
	person.Dream = *dream

	fmt.Printf("Model created: %v\n", *person)
	return person
}

//Edit adalah
func Edit(person *model.Person) {
	fmt.Printf("Model Before edit : %v\n", *person)

	afterPerson := &model.Person{}

	afterPerson = &model.Person{
		ID:      person.ID,
		Name:    "Indra",
		Age:     person.Age,
		Address: "Malang",
		Dream:   person.Dream,
	}

	fmt.Printf("Model After edit : %v\n", *afterPerson)
}
