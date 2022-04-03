package routes

import (
  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/help"
  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
)

func ViewUser(c *gin.Context) {
  username := c.Param("username")

  var u models.User
  result := db.DB.Where("username = ?", username).First(&u)

  if result.RowsAffected == 0 {
    help.APIResponse(c, 404, "NOTFOUND", username)
    return
  }

  data := map[string]interface{}{
    "user": u,
  }

  help.View(c, "viewUser", "main", data)
}

func ViewUserArticles(c *gin.Context) {
  userID := c.Param("id")

  data := map[string]interface{}{
    "userID": userID,
  }

  help.View(c, "viewUserArticles", "main", data)
}

func Profile(c *gin.Context) {
  var data map[string]interface{}

  help.View(c, "profile", "main", data)
}
