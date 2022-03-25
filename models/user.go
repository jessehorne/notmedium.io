package models

import (
  "time"

  "gorm.io/gorm"
)

type User struct {
  gorm.Model

  Username string `gorm:"type:varchar(255);unique"`
  Password string `gorm:"type:varchar(255)"`
  ApiToken string `gorm:"type:varchar(255)"`
  ApiTokenExpiresAt time.Time
  IsAdmin bool
}
