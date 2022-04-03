package main

import (
  "strconv"
  // "fmt"

  "github.com/jessehorne/notmedium.io/db"
  "github.com/jessehorne/notmedium.io/models"
)

func main() {
  for x := 1; x <= 30; x++ {
    newArticle := models.Article{
      Author: "dock",
      Title: "beep bop boop number " + strconv.Itoa(x),
      Content: "Here is a story about how my life got flipped turned upside down #" + strconv.Itoa(x),
      Rank: x,
      Published: true,
    }

    db.DB.Create(&newArticle)
  }
}
