package main

import "fmt"

func main()  {
  var nilai32 = 100000
  var nilai64 int64 = int64(nilai32)
  var nilai8 int8 = int8(nilai32)

  fmt.Println(nilai32)
  fmt.Println(nilai64)
  fmt.Println(nilai8)

  var name = "Eko"
  var e byte = name[0]
  var estring string = string(e)
  var k byte = name[1]
  kstring := string(k)

  fmt.Println(name)
  fmt.Println(estring)
  fmt.Println(kstring)
}
