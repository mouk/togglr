# togglr [![Build Status](https://travis-ci.org/mouk/togglr.svg?branch=master)](https://travis-ci.org/mouk/togglr)&nbsp;[![godoc reference](https://godoc.org/github.com/mouk/togglr?status.png)][godoc]
A simple implementation of the Feature Toggles pattern for Golang


## Feature list
- [X] Configure from environment variables
- [X] Configure from JSON file
- [ ] Percentage based configuration
- [ ] Time base configuration


## Getting Started

~~~ go
package main

import "github.com/mouk/togglr"
type MyFeatures struct{
  Feature1 togglr.Feature
  Feature2 togglr.Feature `env:"f2"`
}

func main() {
  features = MyFeatures{}
  togglr.Read(&features)


  if(features.Feature1.IsEnabled()){
    // Do somthing
  }
}
~~~
