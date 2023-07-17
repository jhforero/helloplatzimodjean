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

func (ftEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type PrintInfo interface {
	getMessage() string
}

// func GetMessage(p Person) {
// 	fmt.Printf("%s with age %d\n", p.name, p.age)
// }

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.name = "Jean H Forero M"
	ftEmployee.id = 2
	ftEmployee.age = 30

	// fmt.Print("%v\n", ftEmployee)
	getMessage(ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)

	//GetMessage(ftEmployee.Person)

}
