package main

import "fmt"

type HasName interface{
  GetName() string
}

type Person struct{
  Name string
}

func sayHello(name HasName)  {
  fmt.Println("Hello", name.GetName())
}

func (person Person) GetName() string  {
  return person.Name
}

type Animal struct{
  Name string
}

func (animal Animal) GetName() string  {
  return animal.Name
}



func Ups(i int) interface{} {
  if i == 1 {
    return 1
  } else if i == 2 {
    return true
  } else {
    return "ups"
  }
}

func main()  {
  var eko Person
  eko.Name = "Eko"
  sayHello(eko)

  var cat Animal
  cat.Name = "Pus"
  sayHello(cat)

  var data interface{} = Ups(2)
  fmt.Println(data)
}
