package main

import (
  "fmt"
  "golang-gin-restapi-mysql/db"
  "github.com/gin-gonic/gin"
)

func main() {
  fmt.Println("hello world")
  db.Init()

  router := SetupRouter()
  router.Run(":8081")
}


func SetupRouter() *gin.Engine {
  router := gin.Default()

  v1 := router.Group("api/v1")
  {
    v1.GET("/test", Test)
  }
  return router
}

func Test(c *gin.Context) {
  c.JSON(200, gin.H {"status": 200, "data": "testing api"})
}
