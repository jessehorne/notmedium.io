package middleware

import (
  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/help"
  "github.com/jessehorne/notmedium.io/db"
  "github.com/jessehorne/notmedium.io/models"
)

func Auth(c *gin.Context) {
  // Get API token from Authorization header
  token, exists := c.Request.Header["Authorization"]

  if !exists {
    help.APIAbortResponse(c, 401, "TokenError", "Missing token.")
    return
  }

  // get user according to header
  var user models.User
  result := db.DB.Where("api_token = ?", token).First(&user)

  if result.RowsAffected == 0 {
    help.APIAbortResponse(c, 401, "TokenError", "Invalid token.")
    return
  }

  // store authed user details
  c.Set("user", user)

  c.Next()
}