package main

import (
  "fmt"

  "github.com/gin-gonic/gin"
  "golang-gin-restapi-mysql/controllers"
)

func main() {
  fmt.Println("hello world")

  router := SetupRouter()
  router.Run(":8081")
}


func SetupRouter() *gin.Engine {
  router := gin.Default()

  v1 := router.Group("api/v1")
  {
    v1.POST("/flower", controllers.Create)
    //v1.GET("/flower/:id", controllers.GetFlower)
    v1.GET("/flowers", controllers.GetAllFlower)
    v1.GET("/check", controllers.HealthCheck)
  }
  return router
}
