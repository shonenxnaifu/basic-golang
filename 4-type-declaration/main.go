package main

import "fmt"

func main()  {
  type NoKTP string
  type Married bool

  var noKtpEko NoKTP = "12312313123"
  var marriedStatus Married = true

  fmt.Println(noKtpEko)
  fmt.Println(marriedStatus)
}
