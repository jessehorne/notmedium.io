package db

import (
  "strings"
  "os"

  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "github.com/joho/godotenv"
)

var DB *gorm.DB

func init() {
  err := godotenv.Load()

  if err != nil {
    panic("Missing .env")
  }

  dbName := os.Getenv("DB_NAME")
  dbUser := os.Getenv("DB_USER")
  dbPass := os.Getenv("DB_PASS")

  // setup dsn
  dsn := "USER:PASS@tcp(127.0.0.1:3306)/NAME?charset=utf8mb4&parseTime=True&loc=Local"
  dsn = strings.Replace(dsn, "NAME", dbName, -1)
  dsn = strings.Replace(dsn, "USER", dbUser, -1)
  dsn = strings.Replace(dsn, "PASS", dbPass, -1)

  // setup mysql db
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Error connecting to MySQL database")
  }

  DB = db
}
