package main

import (
	"fmt"
	"shonen-dev.xyz/basic-golang/26-package-initialization/database"
)

func main()  {
  result := database.GetDataBase()
  fmt.Println(result)
}
