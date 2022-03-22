package api

import (
  "strconv"
  "strings"

  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/help"
  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
)

type userUpdateRequest struct {
  Username string `json:"username" binding:"required,max=255,alphanum"`
  DisplayName string `json:"displayName" binding:"required,max=50"`
  Description string `json:"description" binding:"required"`
}

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

func UsersUpdateByID(c *gin.Context) {
  userID := c.Param("id")
  intUserID, _ := strconv.Atoi(userID)

  // get user by id
  var searchUser models.User
  result := db.DB.First(&searchUser, intUserID)

  if result.RowsAffected == 0 {
    help.APIResponse(c, 404, "NotFoundByID", "No user found with that ID.")
    return
  }

  // make sure user is admin OR current user
  isSameUser := searchUser.ID == c.Value("user").(models.User).ID
  isAdmin := c.Value("user").(models.User).IsAdmin

  if !isSameUser && !isAdmin {
    help.APIResponse(c, 401, "PermissionError", "You can't do that.")
    return
  }

  // cast json to struct
  var jsonUser userUpdateRequest
  err := c.BindJSON(&jsonUser)

  if err != nil {
    help.APIResponse(c, 400, "ValidationError", strings.Split(err.Error(), "\n"))
    return
  }

  // only update certain values
  if jsonUser.Username != "" {
    searchUser.Username = jsonUser.Username
  }

  if jsonUser.DisplayName != "" {
    searchUser.DisplayName = jsonUser.DisplayName
  }

  if jsonUser.Description != "" {
    searchUser.Description = jsonUser.Description
  }

  // update user
  db.DB.Save(&searchUser)

  help.APIResponse(c, 200, "OK", searchUser)
}

func UsersDeleteByID(c *gin.Context) {
  userID := c.Param("id")
  intUserID, _ := strconv.Atoi(userID)

  // get user by id
  var searchUser models.User
  result := db.DB.First(&searchUser, intUserID)

  if result.RowsAffected == 0 {
    help.APIResponse(c, 404, "NotFoundByID", "No user found with that ID.")
    return
  }

  // make sure user is admin OR current user
  isSameUser := searchUser.ID == c.Value("user").(models.User).ID
  isAdmin := c.Value("user").(models.User).IsAdmin

  if !isSameUser && !isAdmin {
    help.APIResponse(c, 401, "PermissionError", "You can't do that.")
    return
  }

  // update user
  db.DB.Delete(&searchUser)

  help.APIResponse(c, 200, "OK", nil)
}
