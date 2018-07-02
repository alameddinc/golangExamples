package main

import "fmt"

func incNum(p *int,n int)  {
  *p+=n
}

func main()  {
  i:=10
  var p *int = &i

  fmt.Println(i)
  fmt.Println(*p)

  incNum(p,20)

  fmt.Println(i)
  fmt.Println(*p)

}
