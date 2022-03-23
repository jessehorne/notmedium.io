package routes

import (
  "github.com/kataras/blocks"
)

var Blocks *blocks.Blocks

func init() {
  Blocks = blocks.New("./views")

  err := Blocks.Load()

  if err != nil {
    panic("Could not load views...")
  }
}
