package models

type Result struct {
  Runner *Runner
  SwimTime  uint64 // swimming time in seconds
  Trasition1Time uint64 // transition time Swim-to-Bike
  BikeTime  uint64 // biking time in seconds
  Trasition2Time uint64 // transition time Bike-to-Run
  RunTime   uint64 // running time in seconds
}
