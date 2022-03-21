package models

import (
  "time"

  "gorm.io/gorm"
)

type User struct {
  gorm.Model

  Email string `gorm:"type:varchar(255);unique"`
  Username string `gorm:"type:varchar(255);unique_index"`
  Password string `gorm:"type:varchar(255)"`
  DisplayName string `gorm:"type:varchar(50)"`
  Description string
  profileImgPath string
  ApiToken string `gorm:"type:varchar(255)"`
  ApiTokenExpiresAt time.Time
  IsAdmin bool
}
