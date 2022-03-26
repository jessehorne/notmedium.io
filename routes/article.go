package routes

import (
  "strconv"

  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
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

func EditArticle(c *gin.Context) {
  articleID := c.Param("id")

  intArticleID, _ := strconv.Atoi(articleID)

  var article models.Article
  result := db.DB.First(&article, intArticleID)

  if result.RowsAffected == 0 {
    c.Redirect(300, "/")
    return
  }

  data := map[string]interface{}{
    "articleID": articleID,
    "title": article.Title,
    "content": article.Content,
  }

  page, err := Blocks.ParseTemplate("editArticle", "main", data)

  if err != nil {
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(page))
}
