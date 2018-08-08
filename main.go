package main

import (
  "fmt"
  "golang-gin-restapi-mysql/db"
)

func main() {
  fmt.Println("hello world")
  db.Init()
}
