package routes

import (
  "github.com/gin-gonic/gin"
)

func NewArticle(c *gin.Context) {
  var data map[string]interface{}

  page, err := Blocks.ParseTemplate("newArticle", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}

func ViewArticle(c *gin.Context) {
  articleID := c.Param("id")

  data := map[string]interface{}{
    "articleID": articleID,
  }

  page, err := Blocks.ParseTemplate("viewArticle", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}
