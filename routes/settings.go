package routes

import (
  "github.com/gin-gonic/gin"
)

func Settings(c *gin.Context) {
  var data map[string]interface{}

  page, err := Blocks.ParseTemplate("settings", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}
