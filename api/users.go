package api

import (
  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/help"
  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
)

func UsersGetAll(c *gin.Context) {
  // pagination
  limit, page := help.GetPaginationDetails(c)

  // get users
  var users []models.User
  result := db.DB.Offset(page).Limit(limit).Find(&users)

  help.APIResponse(c, 200, "OK", &gin.H{
    "page": page,
    "limit": limit,
    "count": result.RowsAffected,
    "users": users,
  })
}
