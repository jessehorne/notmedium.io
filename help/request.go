package help

import (
  "strconv"

  "github.com/gin-gonic/gin"
)

func GetPaginationDetails(c *gin.Context) (int, int) {
  limit, _ := strconv.Atoi(c.Query("limit"))
  page, _ := strconv.Atoi(c.Query("page"))

  // default values
  if limit == 0 {
    limit = 10
  }

  return limit, page
}
