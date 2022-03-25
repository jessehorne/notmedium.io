package routes

import (
  "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
  var data interface{}

  page, err := Blocks.ParseTemplate("register", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}

func Login(c *gin.Context) {
  var data interface{}

  page, err := Blocks.ParseTemplate("login", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}
