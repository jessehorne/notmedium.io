package routes

import (
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/sessions"

  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"

  "github.com/jessehorne/notmedium.io/help"
)

func NewArticle(c *gin.Context) {
  var data map[string]interface{}

  help.View(c, "newArticle", "main", data)
}

func NewArticlePost(c *gin.Context) {
  title := c.PostForm("title")
  content := c.PostForm("content")
  publish := c.PostForm("publish")

  // validate title
  if len(title) < 1 {
    data := map[string]interface{}{
      "error": "The title must be at least 1 character long.",
    }

    help.View(c, "submit", "main", data)
    return
  }

  // validate content
  if len(content) < 1 {
    data := map[string]interface{}{
      "error": "The content must be at least 1 character long.",
    }

    help.View(c, "submit", "main", data)
    return
  }

  session := sessions.Default(c)
  userID := session.Get("userID")
  username := session.Get("username")

  // create article
  newArticle := models.Article{
    Published: false,
  }

  // set other values
  newArticle.Author = username.(string)
  newArticle.UserID = userID.(uint)
  newArticle.Title = title
  newArticle.Content = content

  if publish == "on" {
    newArticle.Published = true
  } else {
    newArticle.Published = false
  }

  db.DB.Save(&newArticle)

  c.Redirect(302, "/a/" + strconv.Itoa(int(newArticle.ID)))
}

func ViewArticle(c *gin.Context) {
  articleID := c.Param("id")

  intArticleID, _ := strconv.Atoi(articleID)

  var article models.Article
  result := db.DB.First(&article, intArticleID)

  // does user own article?
  owned := false

  userIDExists := c.Value("userID")
  var userID uint

  if userIDExists != nil {
    userID = userIDExists.(uint)
  }

  if article.UserID == userID {
    owned = true
  }
  // end does user own

  data := map[string]interface{}{
    "article": article,
    "owned": owned,
  }

  if result.RowsAffected == 0 {
    c.Redirect(300, "/")
    return
  }

  help.View(c, "viewArticle", "main", data)
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

  help.View(c, "editArticle", "main", data)
}

func EditArticlePost(c *gin.Context) {
  articleID := c.Param("id")

  title := c.PostForm("title")
  content := c.PostForm("content")
  publish := c.PostForm("publish")

  var article models.Article
  result := db.DB.First(&article, articleID)

  if result.RowsAffected == 0 {
    c.Redirect(302, "/")
    return
  }

  // validate title
  if len(title) < 1 {
    data := map[string]interface{}{
      "error": "The title must be at least 1 character long.",
      "articleID": articleID,
      "title": article.Title,
      "content": article.Content,
    }

    help.View(c, "editArticle", "main", data)
    return
  }

  // validate content
  if len(content) < 1 {
    data := map[string]interface{}{
      "error": "The content must be at least 1 character long.",
      "articleID": articleID,
      "title": article.Title,
      "content": article.Content,
    }

    help.View(c, "editArticle", "main", data)
    return
  }

  // set other values
  article.Title = title
  article.Content = content

  if publish == "on" {
    article.Published = true
  } else {
    article.Published = false
  }

  db.DB.Save(&article)

  c.Redirect(302, "/a/" + strconv.Itoa(int(article.ID)))
}

func DeleteArticle(c *gin.Context) {
  articleID := c.Param("id")

  var article models.Article
  result := db.DB.Find(&article, articleID)

  if result.RowsAffected == 0 {
    c.Redirect(302, "/")
    return
  }

  // does user own article?
  owned := false

  userIDExists := c.Value("userID")
  var userID uint

  if userIDExists != nil {
    userID = userIDExists.(uint)
  }

  if article.UserID == userID {
    owned = true
  }
  // end does user own

  if owned {
    db.DB.Delete(&models.Article{}, article.ID)
  }

  c.Redirect(302, "/")
  return
}
