package help

import (
  "github.com/gin-gonic/gin"
)

func APIResponse(c *gin.Context, status int, data interface{}) {
  c.JSON(status, gin.H{
    "data": data,
  })
}
