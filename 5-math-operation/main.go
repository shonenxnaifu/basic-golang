package main

import "fmt"

func main()  {
  var a = 10
  var b = 10
  var c = a + b

  var d = 10
  d += 10

  c++

  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(-c)
}
