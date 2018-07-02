package main

import "fmt"

func main()  {

  x:=1
  doubleUp:=func () {
    x*=2
  }
  fmt.Println(doubleUp())
  fmt.Println(doubleUp())
  fmt.Println(doubleUp())
  fmt.Println(doubleUp())
}
