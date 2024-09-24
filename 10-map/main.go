package main

import "fmt"

func main()  {
  var person map[string]string = map[string]string{
    "name": "Eko",
    "address": "Subang",
  }

  person["title"] = "Programmer"

  fmt.Println(person)
  fmt.Println(person["name"])
  fmt.Println(len(person))


  book := make(map[string]string)
  book["title"] = "Belajar GO Lang"
  book["author"] = "Eko"
  book["ups"] = "salah"

  fmt.Println(book)
  fmt.Println(len(book))
  delete(book, "ups")

  fmt.Println(book)
  fmt.Println(len(book))
}
