package main

import (
  "fmt"
  "io/ioutil")

type Pairs struct{
  Id int
  Name string
}

func main()  {
  bs,err := ioutil.ReadFile("test.gofig")
  if err != nil{
    panic(err)
  }

}
