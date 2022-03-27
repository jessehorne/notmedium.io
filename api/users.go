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

func UsersGetArticles(c *gin.Context) {
  userID := c.Param("id")

  intUserID, _ := strconv.Atoi(userID)

  {
    var user models.User
    result := db.DB.First(&user, intUserID)

    if result.RowsAffected == 0 {
      help.APIResponse(c, 404, "NotFoundByID", "No user found with that ID.")
      return
    }
  }

  // pagination
  limit, page := help.GetPaginationDetails(c)
  count := 0

  // get users
  var articles []models.Article
  var returnArticles []models.Article

  result := db.DB.Where("user_id = ?", intUserID).Order("id desc").Find(&articles)

  // pagination
  startIndex := (page * limit)
  endIndex := (page * limit) + limit
  for i := startIndex; int64(i) < result.RowsAffected; i++ {
    if i >= startIndex && i < endIndex {
      returnArticles = append(returnArticles, articles[i])
      count += 1
    }
  }

  help.APIResponse(c, 200, "OK", &gin.H{
    "page": page,
    "limit": limit,
    "count": count,
    "articles": returnArticles,
    "totalCount": result.RowsAffected,
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

  // pagination
  limit, page := help.GetPaginationDetails(c)

  // get users
  var articles []models.Article
  db.DB.Where("published =?", true).Where("user_id = ?", intUserID).Offset(page).Limit(limit).Find(&articles)

  help.APIResponse(c, 200, "OK", gin.H{
    "user": user,
    "articles": articles,
  })
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
