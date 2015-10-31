# togglr [![Build Status](https://travis-ci.org/mouk/togglr.svg?branch=master)](https://travis-ci.org/mouk/togglr)&nbsp;[![GoDoc](https://godoc.org/github.com/mouk/togglr?status.svg)](http://godoc.org/github.com/mouk/togglr)
A simple implementation of the Feature Toggles pattern for Golang


## Feature list
- [X] Configure from environment variables
- [X] Configure from JSON file
- [X] Randomly enabled feature
- [X] Time base configuration
- [X] Percentage based configuration
- [X] Week day based configuration


## Getting Started

~~~ go
package main

import "github.com/mouk/togglr"
type MyFeatures struct{
  Feature1 togglr.Feature
  Feature2 togglr.Feature `env:"f2"`
}

func main() {
  features := MyFeatures{}
  togglr.Read(&features)


  if(features.Feature1.IsEnabled()){
    // Do somthing
  }
}
~~~


## Export feature snapshot
A static overview of all features can be exported as a dictionary.
This should be used only for debugging and reporting puposes,
but not to change the control flow.

~~~ go
package main

import "github.com/mouk/togglr"

type MyFeatures struct{
  Feature1 togglr.Feature
  //....
}

func main() {
  togglr.Init("features.json")
  features := MyFeatures{}
  togglr.Read(&features)

  snapshot = togglr.GetFeatureSpanshots(&features)

  for feature, enabled := snapshot m {
    fmt.Println("Feature:", feature, "Enabled:", enabled)
  }
}
~~~
