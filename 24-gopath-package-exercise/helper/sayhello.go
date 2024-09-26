package helper

import "fmt"

var version = "1.0.0"
var Application = "Belajar Golang"

func SayHello(name string)  {
  fmt.Println("Hello", name)
}

func sayHello(name string)  {
  fmt.Println("Hello", name)
}
