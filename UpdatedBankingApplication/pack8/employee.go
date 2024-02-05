package pack8

import (
	"fmt"
)

type Employee struct {
	Firstname     string
	Lastname      string
	Age           int
	Experience    string
	Qualification string
	Empid         int
}

func InitialiseEmployee() []Employee {
	emp1 := Employee{
		Firstname:     "Anurag",
		Lastname:      "Dullur",
		Age:           24,
		Experience:    "2years",
		Qualification: "MS",
		Empid:         1,
	}
	emp2 := Employee{
		Firstname:     "Haris",
		Lastname:      "Kolagani",
		Age:           24,
		Experience:    "2years",
		Qualification: "MS",
		Empid:         2,
	}
	emp3 := Employee{
		Firstname:     "Durga Mounika",
		Lastname:      "Jagarlamudi",
		Age:           25,
		Experience:    "3years",
		Qualification: "MS",
		Empid:         3,
	}
	var emp = make([]Employee, 3)
	emp[0] = emp1
	emp[1] = emp2
	emp[2] = emp3
	return emp
}

func LoginAsEmployee(employee []Employee, empid int) error {
	for i := 0; i < len(employee); i++ {
		if employee[i].Empid == empid {
			return nil
		}
	}
	return fmt.Errorf("Employee not found")

}
