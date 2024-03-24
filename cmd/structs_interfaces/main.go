package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name      string
	age       uint8
	jobDetail Job
}

type Job struct {
	jobName     string
	description string
}

func (e Job) getName() string {
	return e.jobName
}

// we can also attach methods to structs like this
func (e Employee) getName() string {
	return e.name
}

//the method below takes in an Employee type and just returns the name in UpperCase
// currently we can't pass any other type to this function except employee
// This is where interfaces come in , we can define an interface and have methods defined within it which are implemented by various types

// func nameToUppercase(e Employee) string{
// return strings.ToUpper(e.getName())
// }

func nameToUppercase(e genericInterface) string {
	return strings.ToUpper(e.getName())
}

type genericInterface interface {
	getName() string
}

func main() {
	var jobDetail = Job{"Software Engineer", "Staring at a screen all day trying to find that null pointer exception"}
	var employee1 = Employee{"shadaan", 27, jobDetail}
	fmt.Println(employee1)
	//we can also define anonymous structs , but it has to be declared and initialised at the same time
	var employee2 = struct {
		name string
		age  uint8
	}{"shafia", 29}
	fmt.Println(employee2)
	fmt.Println(employee1.getName())

	var name1 = nameToUppercase(employee1)
	var name2 = nameToUppercase(jobDetail)

	fmt.Printf("Used interfaces and got names converted to upperacase : %v %v", name1, name2)
}
