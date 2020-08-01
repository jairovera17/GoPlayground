package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//alex := person{"Alex", "Andreson"}
	alex := person{firstName: "Alex", lastName: "Andreson"}
	fmt.Println(alex)
}
