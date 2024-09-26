package main

import (
	"flag"
	"fmt"
)

func main()  {
  host := flag.String("host", "localhost", "Put your database host")
  username := flag.String("username", "username", "Put your database username")
  password := flag.String("password", "pass", "Put your database password")
  var number *int = flag.Int("number", 100, "Put your number")

  flag.Parse()

  fmt.Println(*host, *username, *password)
  fmt.Println("Number: ", *number)
}
