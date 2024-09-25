package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello()  {
  fmt.Println("Hello ", customer.Name)
}

func main() {
  var eko Customer
  eko.Name = "Eko"
  eko.Address = "indonesia"
  eko.Age = 30

  fmt.Println(eko.Name, eko.Address, eko.Age)

  joko := Customer{
    Name: "joko",
    Address: "indo",
    Age: 39,
  }
  fmt.Println(joko)

  budi := Customer{"Budi", "jkt", 2}
  fmt.Println(budi)
  budi.sayHello()
}
