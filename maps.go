package main

import "fmt"

func main()  {
  student:= make(map[string]int)
  student["ali"]=78
  student["veli"]=84
  student["kerem"]=58
  student["ayse"]=87
  student["asli"]=85

  fmt.Println(student["ayse"])
}
