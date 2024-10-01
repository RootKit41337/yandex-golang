package main

import (
	"fmt"
)

type Employee struct{
	name string
	position string
	salary float64
	bonus float64
}

func (e Employee) CalculateTotalSalary(){
	fmt.Printf("Employee: %s\n", e.name)
	fmt.Printf("Position: %s\n", e.position)
	fmt.Printf("Total Salary: %.2f", e.salary + e.bonus)
}