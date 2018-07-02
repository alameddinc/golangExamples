package main

import "fmt"

func main()  {
  i:=10
  var p *int = &i

  fmt.Println(i)
  fmt.Println(*p)
  fmt.Println(p)
  fmt.Println(&i)

  //Pointer's Pointer
  var p2 **int = &p
  fmt.Println(p2)
  fmt.Println(*p2)
  fmt.Println(**p2)



}
