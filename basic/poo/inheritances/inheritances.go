package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	endData string
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporal Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(print PrintInfo) {
	fmt.Println(print.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.id = 1
	ftEmployee.name = "Raquel"
	ftEmployee.age = 32

	fmt.Printf("%v\n", ftEmployee)

	tEmployee := TemporaryEmployee{}

	getMessage(tEmployee)
	getMessage(ftEmployee)

}
