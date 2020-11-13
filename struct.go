package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	age       int
}

func (u user) detail() {
	fmt.Println(u.firstName, u.lastName, u.age)
}
func main() {
	x := user{age: 30, firstName: "John", lastName: "Anderson"}
	x.detail()
}
