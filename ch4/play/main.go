package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func EmployeeByID(id int) *Employee { /* */
	return &Employee{}
}

func main() {

	var dilbert Employee

	dilbert.Salary = 100000
	position := "Software Engineer"
	dilbert.Position = position

	{
		position := &dilbert.Position
		*position = "Senior " + *position
	}

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // PHB

	id := dilbert.ID
	EmployeeByID(id).Salary = 0

	fmt.Println(employeeOfTheMonth.Position)
}
