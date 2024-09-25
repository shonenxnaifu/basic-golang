package main

import "fmt"

type Address struct{
  City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address)  {
  address.Country = "Belanda"
}

type Man struct{
  Name string
}

func (man *Man) Married()  {
  man.Name = "Mr. " + man.Name
}

func main()  {
  address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
  address2 := &address1
  var address3 *Address = &address1 
  var address4 *Address = &Address{"Surabaya", "Jawa Timur", "Indonesia"}


  address2.City = "Bandung"

  *address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

  fmt.Println(address1)
  fmt.Println(address2)
  fmt.Println(address3)
  fmt.Println(address4)

  var address5 *Address = new(Address)
  address5.City = "Jakarta"
  fmt.Println(address5)

  var alamat = Address{
    City: "Subang",
    Province: "Jawa Barat",
    Country: "Indonesia",
  }

  var alamatPointer *Address = &alamat

  //ChangeCountryToIndonesia(&alamat)
  ChangeCountryToIndonesia(alamatPointer)
  fmt.Println(alamat)
  fmt.Println(alamatPointer)

  eko := Man{"Eko"}
  eko.Married()

  fmt.Println(eko.Name)
}
