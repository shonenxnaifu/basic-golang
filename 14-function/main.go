package main

import "fmt"

type Filter func(string) string
type BlackList func(string) bool

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string, string) {
	return "Eko", "Khannedy", "Lorem"
}

func getFullNameReturnValues() (firstName, lastName string) {
	firstName = "Eko"
	lastName = "Kurniawan"
	return
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", filter(nameFiltered))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	sayHello()
	sayHelloTo("Eko", "Kurniawan")
	result := getHello("Budi")
	fmt.Println(result)

	firstName, lastName, _ := getFullName()
	// firstName, _ := getFullName()
	fmt.Println(firstName, lastName)

	a, b := getFullNameReturnValues()
	fmt.Println(a, b)

	total := sumAll(10, 20, 30, 40)
	fmt.Println(total)

	numbers := []int{10, 20, 40, 50, 20}
	total = sumAll(numbers...)
	fmt.Println(total)

	sayHelloWithFilter("Anjing", spamFilter)

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	factorialResult := factorialRecursive(5)
	fmt.Println(factorialResult)
}
