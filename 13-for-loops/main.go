package main

import "fmt"

func main()  {
  //counter := 1
  //for counter <= 10 {
  //  fmt.Println("Perulangan ke ", counter)
  //  counter++
  //}

  for counter := 1; counter <=10; counter++ {
    fmt.Println("Perulangan ke ", counter)
  }

  slice := []string{"Eko", "Kurniawan", "Khannedy"}

  for i := 0; i < len(slice); i++ {
    fmt.Println(slice[i])
  }

  for index, value := range slice {
    fmt.Println("Index", index, "=", value)
  }

  for i := 0; i < 10; i++ {
    //if i == 5{
    //  break
    //}

    if i % 2 == 0 {
      continue
    }
    fmt.Println("Perulangan ke", i)
  }
}
