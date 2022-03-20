package main

import (
  "fmt"

  "github.com/jessehorne/notmedium.io/db"
  "github.com/jessehorne/notmedium.io/models"
)

func main() {
  fmt.Println("Migrating...")

  db.DB.AutoMigrate(&models.User{})
  db.DB.AutoMigrate(&models.Article{})
  db.DB.AutoMigrate(&models.Tag{})
  db.DB.AutoMigrate(&models.ArticleTag{})

  fmt.Println("Done!")
}
