package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error)  {
  if pembagi == 0 {
    return 0, errors.New("Pembagi tidak boleh nol")
  } else {
    result := nilai / pembagi
    return result, nil
  }
}

func main()  {
  var contohError error = errors.New("Ups Error")
  fmt.Println(contohError.Error())

  hasil, err := Pembagi(20, 2)
  if err == nil {
    fmt.Println("hasil: ", hasil)
  } else {
    fmt.Println("Error: ", err)
  }
}
