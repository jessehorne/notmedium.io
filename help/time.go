package help

import (
  "time"
  "fmt"
)

func GetLaterTime(days int) time.Time {
  return time.Now().Local().Add(time.Minute * 60 * 24 * time.Duration(days))
}

func GetAgo(t time.Time) string {
  year, month, day := t.Date()
  year2, month2, day2 := time.Now().Date()

  years := year2 - year
  months := int(month2) - int(month)
  days := day2 - day

  totalMins := int(time.Now().Sub(t).Minutes())
  hours := int(totalMins / 60)
  mins := int(totalMins - (hours * 60))

  var ago string

  if years > 0 {
    ago = fmt.Sprintf("%d years", years)

    if months > 0 {
      ago = fmt.Sprintf("%s %d", ago, months)
    }
  } else {
    if months > 0 {
      ago = fmt.Sprintf("%d months", months)
    } else {
      if days > 0 {
        ago = fmt.Sprintf("%d days", days)
      } else {
        if hours > 0 {
          ago = fmt.Sprintf("%d hours", hours)
        } else {
          if mins > 0 {
            ago = fmt.Sprintf("%d mins", mins)
          } else {
            ago = "seconds"
          }
        }
      }
    }
  }

  ago = fmt.Sprintf("%s ago", ago)

  return ago
}
