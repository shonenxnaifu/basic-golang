package main

import "fmt"

func main()  {
  name  := "Eko"
  counter := 0

  // variable yang diinisialisasi lg dalam block function
  // akan menjadi variable baru pada block scope function tersebut
  // sedangkan variable yang diakses dari dalam block function bisa diubah nilainya
  increment := func ()  {
    name := "Budi"
    fmt.Println("Increment")
    counter++
    fmt.Println(name)
  }

  increment()
  increment()

  fmt.Println(counter)
  fmt.Println(name)
}
