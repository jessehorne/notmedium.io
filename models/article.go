package models

import (
  "gorm.io/gorm"
)

type Article struct {
  gorm.Model

  UserID uint
  Author string

  Title string `gorm:"type:varchar(255)"`
  Content string `gorm:"type:text"`
  Published bool

  Rank int
}

type ArticleResponse struct {
  ID uint
  Author string
  Title string `gorm:"type:varchar(255)"`
  Published bool
  Rank int

  CreatedAgo string
}
