package main

import "fmt"
  //Caesar Crypto
func fibonacci (num int) int{
  if num<=1{
    return 1
  }
  return fibonacci(num-1)+fibonacci(num-2)
}



func main()  {
  for i := 0; i < 15; i++ {
    fmt.Println(fibonacci(i))
  }
}
