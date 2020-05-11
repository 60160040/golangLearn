package main

import "fmt"

func main() {
	a := new(Android)
	a.Talk() //OR a.Person.Talk()
}

// Android is Person
type Android struct {
	Person //OR Person Person
	Model  string
}

// Person is person
type Person struct {
	Name string
}

//Talk is method Print text
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
