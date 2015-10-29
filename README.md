# togglr [![Build Status](https://travis-ci.org/mouk/togglr.svg?branch=master)](https://travis-ci.org/mouk/togglr)&nbsp;[![Coverage Status](https://coveralls.io/repos/mouk/togglr/badge.svg?branch=master&service=github)](https://coveralls.io/github/mouk/togglr?branch=master)&nbsp;[![GoDoc](https://godoc.org/github.com/mouk/togglr?status.svg)](http://godoc.org/github.com/mouk/togglr)
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
