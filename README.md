# togglr
A simple implementation of the Feature Toggles pattern for Golang

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
