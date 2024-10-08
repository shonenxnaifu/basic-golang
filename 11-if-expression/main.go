package main

import "fmt"

func main()  {
  name := "Joko Susoyo"

  if name == "Eko" {
    fmt.Println("Hello Eko")
  } else if name == "Joko" {
    fmt.Println("Hello Joko")
  } else {
    fmt.Println("Gak hello")
  }

  // inisialisasi variable sekaligus dijadikan kondisi
  if length := len(name); length > 5 {
    fmt.Println("Terlalu panjang")
  } else {
    fmt.Println("Sudah benar")
  }
}
