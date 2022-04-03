package help

import (
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/kataras/blocks"

  "github.com/gin-contrib/sessions"
)

var Blocks *blocks.Blocks

func init() {
  Blocks = blocks.New("./views")

  err := Blocks.Load()

  if err != nil {
    panic("Could not load views...")
  }
}

func GetPaginationDetails(c *gin.Context) (int, int) {
  limit, _ := strconv.Atoi(c.Query("limit"))
  page, _ := strconv.Atoi(c.Query("page"))

  // default values
  if limit == 0 {
    limit = 10
  }

  return limit, page
}

func View(c *gin.Context, view string, layout string, data interface{}) {
  session := sessions.Default(c)
  authed := session.Get("authed")

  newData := map[string]interface{}{
    "data": data,
    "authed": authed,
  }

  html, err := Blocks.ParseTemplate(view, layout, newData)

  if err != nil {
    APIResponse(c, 500, "TemplateParseError", err.Error())
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(html))
}
