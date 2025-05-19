package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 24

	return &p
}

func learnStruct() {

	fmt.Println(person{"Mohamed", 24})
	fmt.Println(person{"Saber", 19})
	fmt.Println(person{name: "Mohamed"})
	fmt.Println(&person{"Mohamed", 24})
	fmt.Println(newPerson("John"))

	s := person{"Ben Malango", 31}

	sp := &s

	fmt.Println(sp.age)
}
