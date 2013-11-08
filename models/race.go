package models

import (
  "time"
)

type Race struct {
  Key string
  Name string
  Date time.Time
}
