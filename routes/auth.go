package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/sessions"

  "github.com/jessehorne/notmedium.io/help"
  "github.com/jessehorne/notmedium.io/models"
  "github.com/jessehorne/notmedium.io/db"
)

func Register(c *gin.Context) {
  var data interface{}

  help.View(c, "register", "main", data)
}

func Login(c *gin.Context) {
  var data interface{}

  help.View(c, "login", "main", data)
}

func RegisterPost(c *gin.Context) {
  username := c.PostForm("username")
  password := c.PostForm("password")
  passwordConfirm := c.PostForm("passwordConfirm")

  // validate username (exists, unique)
  if username == "" {
    data := map[string]interface{}{
      "error": "You must include a username.",
    }

    help.View(c, "register", "main", data)
    return
  }

  var foundUser models.User
  result := db.DB.Where("username =?", username).First(&foundUser)

  if result.RowsAffected != 0 {
    data := map[string]interface{}{
      "error": "That user already exists!",
    }

    help.View(c, "register", "main", data)
    return
  }

  // validate password (min 8, matches passwordConfirm)
  if password != passwordConfirm {
    data := map[string]interface{}{
      "error": "Your password and password confirmation should match.",
    }

    help.View(c, "register", "main", data)
    return
  }

  if len(password) < 4 {
    data := map[string]interface{}{
      "error": "Password must be at least 4 characters long.",
    }

    help.View(c, "register", "main", data)
    return
  }

  // create user

  // generate password
  hashedPassword := help.HashPassword(password)

  apiToken := help.GenerateApiToken()

  // create user struct
  newUser := models.User{
    Username: username,
    Password: hashedPassword,
    ApiToken: apiToken,
    ApiTokenExpiresAt: help.GetLaterTime(1),
  }

  // attempt to create record in db
  createdUser := db.DB.Create(&newUser)

  if createdUser.Error != nil {
    data := map[string]interface{}{
      "error": createdUser.Error.Error(),
    }

    help.View(c, "register", "main", data)
    return
  }

  data := map[string]interface{}{
    "message": "You've successfully registered! You can now log in.",
  }

  help.View(c, "login", "main", data)
}

func LoginPost(c *gin.Context) {
  // data
  username := c.PostForm("username")
  password := c.PostForm("password")

  // validate user
  var foundUser models.User
  result := db.DB.Where("username =?", username).First(&foundUser)

  if result.RowsAffected == 0 {
    data := map[string]interface{}{
      "error": "That user doesn't exist.",
    }

    help.View(c, "login", "main", data)
    return
  }

  // validate password
  // user found, check password
  if !help.CheckPassword(password, foundUser.Password) {
    data := map[string]interface{}{
      "error": "Invalid password!",
    }

    help.View(c, "login", "main", data)
    return
  }

  // handle session
  session := sessions.Default(c)
  session.Set("username", foundUser.Username)
  session.Set("userID", foundUser.ID)
  session.Set("authed", true)
  session.Save()

  c.Redirect(302, "/")
}

func Logout(c *gin.Context) {
  // handle session
  session := sessions.Default(c)
  session.Set("username", "")
  session.Set("userID", "")
  session.Set("authed", false)
  session.Save()

  c.Redirect(302, "/login")
}
