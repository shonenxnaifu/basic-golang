package main

import "fmt"

func main()  {
  var months = [...]string{
    "Januari",
    "Februari",
    "Maret",
    "April",
    "Mei",
    "Juni",
    "Juli",
    "Agustus",
    "September",
    "Oktober",
    "November",
    "Desember",
  }

  var slice1 = months[4:7]

  fmt.Println(slice1)
  fmt.Println(len(slice1))
  fmt.Println(cap(slice1))

  // months[5] = "diubah"

  // slice1[0] = "Mei Lagi"
  // fmt.Println(months)

  var slice2 = months[11:]
  fmt.Println(slice2)

  var slice3 = append(slice2, "Eko")
  fmt.Println(slice3)

  slice3[1] = "Bukan Desember"
  fmt.Println(slice3)

  fmt.Println(slice2)
  fmt.Println(months)

}
