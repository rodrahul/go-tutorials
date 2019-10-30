package main

import "go-tutorials/oop_golangbot/employee"

func main() {
	var e employee.Employee
	e = employee.Employee{
		FirstName:   "Rahul",
		LastName:    "Rode",
		TotalLeaves: 10,
		LeavesTaken: 4,
	}

	e.LeavesRemaining()
}
