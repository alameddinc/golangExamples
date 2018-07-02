package main

import "fmt"

func counter() func() int {
    i:=int(0)
    //"ret int" is important because. if you want to return a function in a function.
    // You need to define inner function's return data  first
    // and its by defined only equal symbol "="
    // because its not a variable its only return data
    return func () (ret int) {
      i++
      ret = i // if it before i++ -> 0 1 2 3 4 5 ...
      return
    }
  }

func main()  {
  test := counter()
  fmt.Println(test())
  fmt.Println(test())
  fmt.Println(test())
  fmt.Println(test())

}
