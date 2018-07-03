package main

import (
  "fmt"
  "math"
)
type geometry interface {
    area() float64
}

type square struct {
  side float64
}
type circle struct {
  radius float64
}

func (c *circle)area() float64 {
  return math.Pi*math.Pow(c.radius,2)
}
func (s *square)area() float64{
  return math.Pow(s.side,2)
}

func result(g geometry) {
    fmt.Println(g.area())
}


func main()  {
  s:=square{side:10.2}
  c:=circle{radius:10}

  result(&c)
  result(&s)
}
