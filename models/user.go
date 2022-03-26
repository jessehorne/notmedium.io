package models

import (
  "time"
  "errors"

  "gorm.io/gorm"

  "github.com/jessehorne/notmedium.io/db"
)

type User struct {
  gorm.Model

  Username string `gorm:"type:varchar(255);unique"`
  Password string `gorm:"type:varchar(255)"`
  ApiToken string `gorm:"type:varchar(255)"`
  ApiTokenExpiresAt time.Time
  IsAdmin bool
}

func GetUsernameByID(id uint) (string, error) {
  var user User
  result := db.DB.Find(&user, id)

  if result.RowsAffected == 0 {
    return "", errors.New("No such user found.")
  }

  return user.Username, nil
}
