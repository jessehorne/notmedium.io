package routes

import (
  "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
  var data interface{}

  page, err := Blocks.ParseTemplate("index", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}
