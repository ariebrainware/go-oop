// This file contain Struct and Method example
// struct: type StructName struct
// method: func (r ReceiverType) funcName(parameters) (results)
// interface: type interfaceName interface {}
package main

import (
	"fmt"
	"time"
)

// This is how we declare a struct
type UserDetail struct {
	firstName string
	lastName  string
	age       int
}

// Here the structs takes UserDetail struct as a field
type JobProfile struct {
	UserDetail
	profile string
}

// Data employee
type DataEmployee struct {
	JobProfile
	salary int
}

type Birthday interface {
	determineBornYear() int
}

type calculateYearOfBirth struct {
	age int
}

func (e DataEmployee) displayYearlySalary() int {
	salaryPerYear := e.salary * 12
	return salaryPerYear
}

func (e *DataEmployee) updateSalary(newSalary int) int {
	e.salary = newSalary
	return e.salary
}

func convertSGDtoIDR(salary int) {
	fmt.Println("Your salary in IDR: ", salary*10482)
}

func (e *UserDetail) changeFirstName(firstName string) string {
	e.firstName = firstName
	return e.firstName
}

func (userAge calculateYearOfBirth) determineBornYear() int {
	y := time.Now().Year()
	return y - userAge.age
}

func calculate(calcAge Birthday) int {
	res := calcAge.determineBornYear()
	return res
}

func main() {
	//Type-1 to define a struct.
	var user1 = UserDetail{"Paul", "J", 19}

	fmt.Println(user1)
	//Type-2 to define a struct
	var user2 = UserDetail{
		firstName: "J",
		lastName:  "Smith",
		age:       20,
	}

	fmt.Println(user2)

	//This gives a specific value from the struct.
	fmt.Println(user2.firstName)

	// We can use the value of a struct from any other struct using the following property
	var user3 = UserDetail{
		firstName: user1.firstName,
		lastName:  user2.lastName,
		age:       user2.age,
	}
	fmt.Println(user3)

	//Structs are immutable
	user3.age = 25
	fmt.Println(user3)

	applicant := JobProfile{UserDetail{"Timmy", "Smith", 19}, "Software Developer"}
	fmt.Println(applicant)

	DataEmploye1 := DataEmployee{
		JobProfile{
			UserDetail{
				user1.firstName,
				user1.lastName,
				user1.age,
			},
			"Software Engineer",
		},
		650,
	}
	fmt.Println("Salary per year: ", DataEmploye1.displayYearlySalary())
	salaryPerYear := DataEmploye1.displayYearlySalary()
	convertSGDtoIDR(salaryPerYear)
	updateSalary := DataEmploye1.updateSalary(800)
	fmt.Println("Your salary is updated from ", DataEmploye1.salary, ", become: ", updateSalary)

	// Change user3 firstname
	user3.changeFirstName("Rectulloci")
	fmt.Println(user3)

	// Calculate Year of Birth from user1
	var data1 = calculateYearOfBirth{user1.age}
	fmt.Println(user1.firstName, "Born in:", calculate(data1))
}
