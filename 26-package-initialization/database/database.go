package database

import "fmt"

var connection string

func init()  {
  fmt.Println("init...")
  connection = "MySQL"
}

func GetDataBase() string {
  return connection
}

