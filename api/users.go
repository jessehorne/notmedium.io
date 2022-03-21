package api

import (
  "strconv"

  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/help"
  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
)

func UsersGetAll(c *gin.Context) {
  // requires user to be admin
  if !c.Value("user").(models.User).IsAdmin {
    help.APIResponse(c, 401, "PermissionError", "You can't do that.")
    return
  }
  
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

func UsersGetOneByID(c *gin.Context) {
  userID := c.Param("id")

  intUserID, _ := strconv.Atoi(userID)

  var user models.User
  result := db.DB.First(&user, intUserID)

  if result.RowsAffected == 0 {
    help.APIResponse(c, 404, "NotFoundByID", "No user found with that ID.")
    return
  }

  help.APIResponse(c, 200, "OK", user)
}
