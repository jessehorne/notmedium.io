package help

import (
  "time"
)

func GetLaterTime(mins int) time.Time {
  return time.Now().Local().Add(time.Minute * time.Duration(mins))
}
