package main

import (
  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/api"
)

var Router *gin.Engine
var Routes *gin.RouterGroup

func init() {
  Router = gin.Default()

  Routes = Router.Group("/api")
  {
    Routes.POST("/register", api.AuthRegister)
    Routes.POST("/login", api.AuthLogin)

    Routes.GET("/users", api.UsersGetAll)
  }
}

func main() {
  Router.Run(":8080")
}
