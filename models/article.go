package models

import (
  "gorm.io/gorm"
)

type Article struct {
  gorm.Model

  Title string `gorm:"type:varchar(255)"`
  ImgPath string `gorm:"type:varchar(255)"`
  Content string `gorm:"type:text"`
  Published bool
  Listed bool
  MetaTitle string `gorm:"type:varchar(255)"`
  MetaDesc string `gorm:"type:varchar(255)"`
  MetaImgPath string `gorm:"type:varchar(255)"`
  Views int64
  Reads int64
  Score int64
}
