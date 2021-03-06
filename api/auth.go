package api

import (
  "strings"

  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
  "github.com/jessehorne/notmedium.io/help"
)

type RegisterRequest struct {
  Username string `json:"username" binding:"required,max=255,alphanum"`
  Password string `json:"password" binding:"required,max=255"`
  PasswordConfirm string `json:"passwordConfirm" binding:"required,max=255"`
}

type LoginRequest struct {
  Username string `json:"username" binding:"required,max=255,alphanum"`
  Password string `json:"password" binding:"required,max=255"`
}

func AuthRegister(c *gin.Context) {
  // validate input
  var req RegisterRequest

  err := c.BindJSON(&req)

  if err != nil {
    help.APIResponse(c, 400, "ValidationError", strings.Split(err.Error(), "\n"))
    return
  }

  // Validate that Password matches PasswordConfirm
  if req.Password != req.PasswordConfirm {
    help.APIResponse(c, 400, "ValidationError", "Passwords must match.")
    return
  }

  // generate password
  hashedPassword := help.HashPassword(req.Password)

  apiToken := help.GenerateApiToken()

  // create user struct
  newUser := models.User{
    Username: req.Username,
    Password: hashedPassword,
    ApiToken: apiToken,
    ApiTokenExpiresAt: help.GetLaterTime(1),
  }

  // attempt to create record in db
  createdUser := db.DB.Create(&newUser)

  if createdUser.Error != nil {
    help.APIResponse(c, 400, "DatabaseError", createdUser.Error.Error())
    return
  }

  // should be fine
  help.APIResponse(c, 200, "OK", nil)
}

func AuthLogin(c *gin.Context) {
  // validate input
  var req LoginRequest

  err := c.BindJSON(&req)

  if err != nil {
    help.APIResponse(c, 400, "ValidationError", strings.Split(err.Error(), "\n"))
    return
  }

  // get user from DB
  user := models.User{}
  result := db.DB.Where("username = ?", req.Username).First(&user)

  if result.RowsAffected == 0 {
    // nothing found
    help.APIResponse(c, 404, "ResourceNotFound", nil)
    return
  }

  // user found, check password
  if !help.CheckPassword(req.Password, user.Password) {
    help.APIResponse(c, 401, "AuthError", "Invalid password.")
    return
  }

  // all good! reup user api token
  token := help.GenerateApiToken()
  expires := help.GetLaterTime(1) // four hours from now

  user.ApiToken = token
  user.ApiTokenExpiresAt = expires

  db.DB.Save(&user)

  help.APIResponse(c, 200, "OK", user)
}
