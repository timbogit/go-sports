package models

import (
  "time"
)

type Race struct {
  Key string
  Name string
  Date time.Time
}

var miamiTime, _  = time.Parse(time.RFC3339, "2013-11-02T15:04:05Z")
var austinTime, _ = time.Parse(time.RFC3339, "2013-11-03T15:04:05Z")
var RaceStore = map[string]*Race{
  "Miami703-2013": &Race{"Miami703-2013", "Miami Half-Triathlon 2013", miamiTime },
  "Austin703-2013": &Race{"Austin703-2013", "Austin Half-Triathlon 2013", austinTime},
}
