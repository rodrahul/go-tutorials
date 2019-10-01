package main

import "gopdf/oop_golangbot/employee"

func main() {
	//e := employee.Employee{
	//	FirstName:   "Rahul",
	//	LastName:    "Rode",
	//	TotalLeaves: 10,
	//	LeavesTaken: 4,
	//}
	var e employee.Employee
	e.LeavesRemaining()
}
