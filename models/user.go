package models

import (
  "time"
  "errors"

  "gorm.io/gorm"

  "github.com/jessehorne/notmedium.io/db"
  "github.com/jessehorne/notmedium.io/help"
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

func (u *User) Articles() []ArticleResponse {
  // get articles
  var articles []Article
  var formatted []ArticleResponse
  result := db.DB.Order("`created_at` DESC").Where("published =?", true).Find(&articles)

  if result.RowsAffected == 0 {
    return formatted
  }

  // do CreatedAgo's
  for _,v := range articles {
    ago := help.GetAgo(v.CreatedAt)

    newFormatted := ArticleResponse{
      ID: v.ID,
      Author: v.Author,
      Title: v.Title,
      Published: v.Published,
      Rank: v.Rank,
      CreatedAgo: ago,
    }

    formatted = append(formatted, newFormatted)
  }

  return formatted
}
