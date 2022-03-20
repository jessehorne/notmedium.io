package help

import (
  "github.com/gin-gonic/gin"
)

func APIResponse(c *gin.Context, status int, msg string, data interface{}) {
  c.JSON(status, gin.H{
    "msg": msg,
    "data": data,
  })
}

func APIAbortResponse(c *gin.Context, status int, msg string, data interface{}) {
  c.AbortWithStatusJSON(status, gin.H{
    "msg": msg,
    "data": data,
  })
}
