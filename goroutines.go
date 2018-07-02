package main

import "fmt"

func getKey(i int)  {
  fmt.Println(i)
}

func main()  {
  for i := 0; i < 200; i++ {
    go getKey(i)
  }
}
