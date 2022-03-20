package help

import (
  "encoding/hex"
  "crypto/md5"
  "crypto/rand"

  "golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
  hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)

  if err != nil {
    return ""
  }

  return string(hash)
}

func CheckPassword(pass string, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

  if err != nil {
    return false
  }

  return true
}

func GenerateApiToken() string {
  t := make([]byte, 32)
  rand.Read(t)

  summed := md5.Sum(t)

  return hex.EncodeToString(summed[:])
}
