package main

import "fmt"

type person struct {
	name string
	age int 
}

func main() {
	fmt.Println(person{"Andrew", 26})
	fmt.Println(person{name: "Olha", age: 24})
	fmt.Println(person{name: "Jack Daniels"})
	fmt.Println(&person{name: "Kevin Levrone", age: 40})

	s := person{name:"Vasiliy", age: 49}
	s.age = 23
	fmt.Println(s.age)
	sp := &s
	fmt.Println(sp.name)
	sp.age = 99
	fmt.Println(sp.age)
	fmt.Println(s.age)
	ssss := s
	ssss.age = 11
	fmt.Println(ssss.age)
	fmt.Println(s.age)
}
