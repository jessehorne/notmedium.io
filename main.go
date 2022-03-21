package main

import (
  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/api"
  "github.com/jessehorne/notmedium.io/middleware"
)

var Router *gin.Engine
var Routes *gin.RouterGroup

func init() {
  Router = gin.Default()

  Router.Use(gin.Recovery())

  Routes = Router.Group("/api")
  {
    Routes.POST("/register", api.AuthRegister)
    Routes.POST("/login", api.AuthLogin)

    Routes.GET("/users", middleware.Auth, api.UsersGetAll)
    Routes.GET("/users/:id", middleware.Auth, api.UsersGetOneByID)
  }
}

func main() {
  Router.Run(":8080")
}
