package help

import (
  "time"
)

func GetLaterTime(days int) time.Time {
  return time.Now().Local().Add(time.Minute * 60 * 24 * time.Duration(days))
}
