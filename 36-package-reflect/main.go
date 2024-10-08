package main

import (
	"fmt"
	"reflect"
)

type Sample struct{
  Name string `required:"true" max:"10"`
  Age int `required:"true" max:"2"`
}

func isValid(data interface{}) bool  {
  t := reflect.TypeOf(data)
  for i := 0; i < t.NumField(); i++{
    field := t.Field(i)
    if field.Tag.Get("required") == "true"{
      value := reflect.ValueOf(data).Field(i).Interface()
      if value == ""{
        return false
      }
    }
  }

  return true
}

func main()  {
  sample := Sample{"Eko", 0}

  var sampleType reflect.Type = reflect.TypeOf(sample)

  fmt.Println(sampleType.NumField())
  fmt.Println(sampleType.Field(1).Name)
  fmt.Println(sampleType.Field(0).Tag.Get("required"))
  fmt.Println(sampleType.Field(0).Tag.Get("max"))
  fmt.Println(sampleType.Field(0).Tag.Get("min"))

  fmt.Println(isValid(sample))
}
