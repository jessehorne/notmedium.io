package api

import (
  "strconv"
  "strings"

  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/help"
  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
)

type articleUpdateRequest struct {
  UserID uint `json:"userID,string" binding:"required"`
  Title string `gorm:"type:varchar(255)" json:"title" binding:"required,min=1"`
  Content string `gorm:"type:text" json:"content" binding:"required,min=1"`
  Published *bool `gorm:"type:bool" json:"published,string" binding:"required"`
}

func ArticlesGetAll(c *gin.Context) {
  // pagination
  limit, page := help.GetPaginationDetails(c)

  // get users
  var articles []models.Article
  result := db.DB.Where("published =?", true).Offset(page*limit).Limit(limit).Find(&articles)

  help.APIResponse(c, 200, "OK", &gin.H{
    "page": page,
    "limit": limit,
    "count": result.RowsAffected,
    "articles": articles,
  })
}

func ArticlesGetOneByID(c *gin.Context) {
  articleID := c.Param("id")

  intArticleID, _ := strconv.Atoi(articleID)

  var article models.Article
  result := db.DB.First(&article, intArticleID)

  if result.RowsAffected == 0 {
    help.APIResponse(c, 404, "NotFoundByID", "No article found with that ID.")
    return
  }

  if article.UserID != c.Value("user").(models.User).ID {
    help.APIResponse(c, 401, "Unauthorized", "You can't do that.")
    return
  }

  help.APIResponse(c, 200, "OK", article)
}

func ArticlesUpdateByID(c *gin.Context) {
  articleID := c.Param("id")
  intArticleID, _ := strconv.Atoi(articleID)

  // get user by id
  var searchArticle models.Article
  result := db.DB.First(&searchArticle, intArticleID)

  if result.RowsAffected == 0 {
    help.APIResponse(c, 404, "NotFoundByID", "No article found with that ID.")
    return
  }

  // make sure that article.UserID == logged in user ID or admin
  isSameUser := searchArticle.UserID == c.Value("user").(models.User).ID
  isAdmin := c.Value("user").(models.User).IsAdmin

  if !isSameUser && !isAdmin {
    help.APIResponse(c, 401, "PermissionError", "You can't do that.")
    return
  }

  // cast json to struct
  var jsonArticle articleUpdateRequest
  err := c.BindJSON(&jsonArticle)

  if err != nil {
    help.APIResponse(c, 400, "ValidationError", strings.Split(err.Error(), "\n"))
    return
  }

  searchArticle.Title = jsonArticle.Title
  searchArticle.Content = jsonArticle.Content

  searchArticle.Published = *jsonArticle.Published

  // update user
  db.DB.Save(&searchArticle)

  help.APIResponse(c, 200, "OK", searchArticle)
}

func ArticlesDeleteByID(c *gin.Context) {
  articleID := c.Param("id")
  intArticleID, _ := strconv.Atoi(articleID)

  // get user by id
  var searchArticle models.Article
  result := db.DB.First(&searchArticle, intArticleID)

  if result.RowsAffected == 0 {
    help.APIResponse(c, 404, "NotFoundByID", "No article found with that ID.")
    return
  }

  // make sure that article.UserID == logged in user ID or admin
  isSameUser := searchArticle.UserID == c.Value("user").(models.User).ID
  isAdmin := c.Value("user").(models.User).IsAdmin

  if !isSameUser && !isAdmin {
    help.APIResponse(c, 401, "PermissionError", "You can't do that.")
    return
  }

  // update user
  db.DB.Delete(&searchArticle)

  help.APIResponse(c, 200, "OK", nil)
}

func ArticlesCreate(c *gin.Context) {
  // bind json to struct
  var validateArticle articleUpdateRequest
  err := c.BindJSON(&validateArticle)

  if err != nil {
    help.APIResponse(c, 400, "ValidationError", strings.Split(err.Error(), "\n"))
    return
  }

  newArticle := models.Article{
    Published: false,
  }

  // set other values
  newArticle.Author = c.Value("user").(models.User).Username
  newArticle.UserID = validateArticle.UserID
  newArticle.Title = validateArticle.Title
  newArticle.Content = validateArticle.Content
  newArticle.Published = *validateArticle.Published // can anyone see this?

  db.DB.Save(&newArticle)

  help.APIResponse(c, 200, "OK", newArticle)
}
