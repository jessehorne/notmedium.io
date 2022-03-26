package main

import (
  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/api"
  "github.com/jessehorne/notmedium.io/routes"
  "github.com/jessehorne/notmedium.io/middleware"
)

var Router *gin.Engine
var Routes *gin.RouterGroup

func init() {
  Router = gin.Default()

  Router.Use(gin.Recovery())

  Router.Static("/public", "./public")

  Routes = Router.Group("/api")
  {
    Routes.POST("/register", api.AuthRegister)
    Routes.POST("/login", api.AuthLogin)

    Routes.GET("/users", middleware.Auth, api.UsersGetAll)
    Routes.GET("/users/:id", middleware.Auth, api.UsersGetOneByID)
    Routes.PUT("/users/:id", middleware.Auth, api.UsersUpdateByID)
    Routes.DELETE("/users/:id", middleware.Auth, api.UsersDeleteByID)
    Routes.GET("/users/:id/articles", api.UsersGetArticles)

    Routes.GET("/articles", api.ArticlesGetAll)
    Routes.GET("/articles/:id", middleware.Auth, api.ArticlesGetOneByID)
    Routes.POST("/articles", middleware.Auth, api.ArticlesCreate)
    Routes.PUT("/articles/:id", middleware.Auth, api.ArticlesUpdateByID)
    Routes.DELETE("/articles/:id", middleware.Auth, api.ArticlesDeleteByID)
  }

  // Views
  Router.GET("/", routes.Index)
  Router.GET("/register", routes.Register)
  Router.GET("/login", routes.Login)
  Router.GET("/profile", routes.Profile)
  Router.GET("/new", routes.NewArticle)
  Router.GET("/a/:id/edit", routes.EditArticle)
  Router.GET("/a/:id", routes.ViewArticle)
  Router.GET("/profile/:id", routes.ViewUser)

}

func main() {
  Router.Run(":8080")
}
