package main

import "fmt"

func main() {
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	var friendName string = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	country := "indonesia"
	fmt.Println(country)

	var (
		firstName = "Eko Kurniawan"
		lastName  = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// constant variable
	const firstNameConst = "Eko"
	const lastNameConst = "Khannedy"
	const value = 1000

	fmt.Println(firstNameConst)
	fmt.Println(lastNameConst)
	fmt.Println(value)

	const (
		firstNameConst2 string = "Eko"
		lastNameConst2         = "Khannedy"
		value2                 = 1000
	)

	fmt.Println(firstNameConst2)
	fmt.Println(lastNameConst2)
	fmt.Println(value2)
}
