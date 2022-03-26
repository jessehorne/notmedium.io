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

func ViewUserArticles(c *gin.Context) {
  userID := c.Param("id")

  data := map[string]interface{}{
    "userID": userID,
  }

  page, err := Blocks.ParseTemplate("viewUserArticles", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}

func Profile(c *gin.Context) {
  var data map[string]interface{}

  page, err := Blocks.ParseTemplate("profile", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}
