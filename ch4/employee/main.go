package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

var data = []Employee{
	{
		ID:        0,
		Name:      "Joshua",
		Address:   "Address 1",
		DoB:       time.Now(),
		Position:  "Hulk",
		Salary:    100000000,
		ManagerID: 0,
	},
	{
		ID:        1,
		Name:      "Richard",
		Address:   "Address 2",
		DoB:       time.Now().Add(1000),
		Position:  "SpiderMan",
		Salary:    100000000,
		ManagerID: 0,
	},
	{
		ID:        2,
		Name:      "Tim",
		Address:   "Address 3",
		DoB:       time.Now(),
		Position:  "Professor X",
		Salary:    1000000000,
		ManagerID: 0,
	},
	{
		ID:        3,
		Name:      "Dom",
		Address:   "Somewhere in Queens",
		DoB:       time.Now().Add(3000),
		Position:  "Black Panther",
		Salary:    100000000,
		ManagerID: 0,
	},
	{
		ID:        4,
		Name:      "Eddy",
		Address:   "1 Balint Dr",
		DoB:       time.Now().Add(4000),
		Position:  "Captain America",
		Salary:    10000000,
		ManagerID: 0,
	},
}

func main() {
	EmployeeByID(1).Name = "RichieRich"
	fmt.Println(data[1])
}

func EmployeeByID(id int) *Employee {
	var name *Employee
	for i := 0; i < len(data); i++ {
		if data[i].ID == id {
			name = &data[i]
		}
	}
	return name
}

func ChangeEmployeeName(e *Employee, newName string) {
	e.Name = newName
}
