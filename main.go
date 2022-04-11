package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"

  "github.com/jessehorne/notmedium.io/api"
  "github.com/jessehorne/notmedium.io/routes"
  "github.com/jessehorne/notmedium.io/middleware"
)

var Router *gin.Engine
var Routes *gin.RouterGroup

func init() {
  Router = gin.Default()

  Router.Use(gin.Recovery())

  store := cookie.NewStore([]byte("secret"))
  Router.Use(sessions.Sessions("sesh", store))

  Router.Static("/public", "./public")

  Routes = Router.Group("/api")
  {
    Routes.POST("/register", api.AuthRegister)
    Routes.POST("/login", api.AuthLogin)

    Routes.GET("/users", middleware.Auth, api.UsersGetAll)
    Routes.GET("/users/:id", middleware.Auth, api.UsersGetOneByID)
    Routes.PUT("/users/:id", middleware.Auth, api.UsersUpdateByID)
    Routes.DELETE("/users/:id", middleware.Auth, api.UsersDeleteByID)
    Routes.GET("/users/:id/articles", middleware.Auth, api.UsersGetArticles)

    Routes.GET("/articles", middleware.Auth, api.ArticlesGetAll)
    Routes.GET("/articles/:id", middleware.Auth, api.ArticlesGetOneByID)
    Routes.POST("/articles", middleware.Auth, api.ArticlesCreate)
    Routes.PUT("/articles/:id", middleware.Auth, api.ArticlesUpdateByID)
    Routes.DELETE("/articles/:id", middleware.Auth, api.ArticlesDeleteByID)
  }

  // Views
  Router.GET("/@:username", routes.ViewUser)
  Router.GET("/", routes.Index)
  Router.GET("/new", routes.IndexNew)

  Router.GET("/register", routes.Register)
  Router.GET("/login", routes.Login)
  Router.POST("/register", routes.RegisterPost)
  Router.POST("/login", routes.LoginPost)

  Router.GET("/profile", routes.Profile)
  Router.GET("/submit", middleware.CookieAuth, routes.NewArticle)
  Router.POST("/submit", middleware.CookieAuth, routes.NewArticlePost)
  Router.GET("/a/:id/edit", routes.EditArticle)
  Router.POST("/a/:id/edit", routes.EditArticlePost)
  Router.GET("/a/:id", routes.ViewArticle)
  Router.GET("/a/:id/delete", routes.DeleteArticle)
  Router.GET("/logout", routes.Logout)

}

func main() {
  Router.Run(":8080")
}
