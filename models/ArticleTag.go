package models

import (
  "gorm.io/gorm"
)

type ArticleTag struct {
  gorm.Model

  ArticleID int64
  TagID int64
}
