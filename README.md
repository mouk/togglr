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

At first create a struct containing the features you want to be able to toggle:

~~~ go
type MyFeatures struct{
  SimpleFeature togglr.Feature
  EnabledOnOrBefore togglr.Feature
  EnabledOnOrAfter togglr.Feature
  FiftyFifty togglr.Feature
  OnlyOnWeekend togglr.Feature
  FeatureFromEnv togglr.Feature `env:"FEATURE"`
}
~~~

Then create a json file configuring thoses feature by thier names:

```json
{
  "SimpleFeature": true,
  "EnabledOnOrBefore": { "enabledOnOrBefore": "2006-01-02 15:04:05" },
  "EnabledOnOrAfter": { "enabledOnOrAfter": "2016-01-02" },
  "FiftyFifty": { "percent": 50 },
  "OnlyOnWeekend" : { "days": "Saturday | Sunday" }
}
```
Almost done. Now just fil the feature sruct and use it

~~~ go
package main

import "github.com/mouk/togglr"

func main() {
  features := MyFeatures{}
  togglr.Init("feature_configuration.json")
  togglr.Read(&features)

  if(features.EnabledOnOrBefore.IsEnabled()){
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
