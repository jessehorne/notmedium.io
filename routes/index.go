package routes

import (
  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
  "github.com/jessehorne/notmedium.io/help"
)

// Shows homepage view with top articles
func Index(c *gin.Context) {
  // pagination
  limit, page := help.GetPaginationDetails(c)

  // get articles
  var articles []models.Article
  var formatted []models.ArticleResponse
  result := db.DB.Order("`rank` DESC").Where("published =?", true).Limit(limit).Offset(page*limit).Find(&articles)

  // do CreatedAgo's
  for _,v := range articles {
    ago := help.GetAgo(v.CreatedAt)

    newFormatted := models.ArticleResponse{
      ID: v.ID,
      Author: v.Author,
      Title: v.Title,
      Published: v.Published,
      Rank: v.Rank,
      CreatedAgo: ago,
    }

    formatted = append(formatted, newFormatted)
  }

  if result.RowsAffected == 0 {
    help.APIResponse(c, 200, "OK", "No articles found")
    return
  }

  data := map[string]interface{}{
    "articles": formatted,
  }

  help.View(c, "index", "main", data)
}

func IndexNew(c *gin.Context) {
  // pagination
  limit, page := help.GetPaginationDetails(c)

  // get articles
  var articles []models.Article
  var formatted []models.ArticleResponse
  result := db.DB.Order("`created_at` DESC").Where("published =?", true).Limit(limit).Offset(page*limit).Find(&articles)

  // do CreatedAgo's
  for _,v := range articles {
    ago := help.GetAgo(v.CreatedAt)

    newFormatted := models.ArticleResponse{
      ID: v.ID,
      Author: v.Author,
      Title: v.Title,
      Published: v.Published,
      Rank: v.Rank,
      CreatedAgo: ago,
    }

    formatted = append(formatted, newFormatted)
  }

  if result.RowsAffected == 0 {
    help.APIResponse(c, 200, "OK", "No articles found")
    return
  }

  data := map[string]interface{}{
    "articles": formatted,
  }

  help.View(c, "index", "main", data)
  // help.APIResponse(c, 200, "OK", data)
}
