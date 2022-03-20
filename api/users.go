package api

import (
  "github.com/gin-gonic/gin"

  "github.com/jessehorne/notmedium.io/help"
)

func UsersGetAll(c *gin.Context) {
  help.APIResponse(c, 200, "OK", []int{})
}
