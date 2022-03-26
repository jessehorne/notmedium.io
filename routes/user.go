package routes

import (
  "github.com/gin-gonic/gin"
)

func ViewUser(c *gin.Context) {
  userID := c.Param("id")

  data := map[string]interface{}{
    "userID": userID,
  }

  page, err := Blocks.ParseTemplate("viewUser", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}
