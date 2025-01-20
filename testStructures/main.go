package main

import (
	"fmt"

	"github.com/Daniil-8bit/GoProjects/magazine"
)

func main() {
	var newUser magazine.Subscriber
	newUser.Name = "John"
	fmt.Println(newUser)

	testUser := magazine.Subscriber{Name: "David", Rate: 3.55, Active: false}
	fmt.Println(testUser)

	testUser1 := magazine.Subscriber{Name: "Ivan"}
	testUser2 := magazine.Subscriber{Rate: 4.15}
	testUser3 := magazine.Subscriber{Name: "Antony", Active: true}

	fmt.Println(testUser1)
	fmt.Println(testUser2)
	fmt.Println(testUser3)

	newEmployee := magazine.Employee{Name: "Cavin", Salary: 500.90}
	fmt.Println(newEmployee)

	newAddress := magazine.Address{Street: "1-st", City: "Chicago", State: "Calarado", PostalCode: "123456"}
	fmt.Println(newAddress)

	/*employeeAddress := magazine.Address{Street: "2-nd", City: "LA", State: "Cansas", PostalCode: "654321"}

	employee1 := magazine.Employee{Name: "Petr", Salary: 3000.45, HomeAddress: employeeAddress}

	fmt.Println(employee1)*/

	employee2 := magazine.Subscriber{Name: "Oleg", Rate: 4.31, Active: true}
	employee2.HomeAddress.City = "Some City"
	employee2.HomeAddress.State = "Oklahoma"
	employee2.HomeAddress.Street = "3-rd"
	employee2.HomeAddress.PostalCode = "983475"

	fmt.Println(employee2)

	employee3 := magazine.Employee{Name: "Harry", Salary: 355.56}
	employee3.City = "NY"
	employee3.Street = "4-th"

	employee3.Address.State = "State"
	employee3.Address.PostalCode = "654283"

	fmt.Println(employee3)
}
