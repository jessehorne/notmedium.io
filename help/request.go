package help

import (
  "strconv"
  "html/template"

  "github.com/gin-gonic/gin"
  "github.com/kataras/blocks"

  "github.com/gin-contrib/sessions"
)

var Blocks *blocks.Blocks

var funcMap = template.FuncMap{
    "inc": func(i int) int {
      return i+1
    },
}

func init() {
  Blocks = blocks.New("./views")
  Blocks.Funcs(funcMap)

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
  username := session.Get("username")
  userID := session.Get("userID")

  newData := map[string]interface{}{
    "data": data,
    "authed": authed,
    "username": username,
    "userID": userID,
  }

  html, err := Blocks.ParseTemplate(view, layout, newData)

  if err != nil {
    APIResponse(c, 500, "TemplateParseError", err.Error())
    return
  }

  c.Data(200, "text/html; charset=utf-8", []byte(html))
}
