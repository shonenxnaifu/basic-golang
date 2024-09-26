package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("salah")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}


	number, err := strconv.ParseInt("1000000", 10, 64)

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 2)
  fmt.Println(value)

  valueInt, err := strconv.Atoi("200000")
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println(err.Error())
	}
}
