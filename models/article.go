package models

import (
  "gorm.io/gorm"
)

type Article struct {
  gorm.Model

  UserID uint

  Title string `gorm:"type:varchar(255)"`
  Content string `gorm:"type:text"`
  Desc string `gorm:"type:varchar(255)"`
  ImgPath string `gorm:"type:varchar(255)"`
  Published bool
  Listed bool
}
