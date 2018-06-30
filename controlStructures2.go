package main

import "fmt"

func main()  {

  //For Loop
  fmt.Println("For Loop")
  for i:=0;i<5;i++{
    fmt.Println(i)
  }

  //While Loop
  fmt.Println("While Loop (if you enter -1, the loop will stop)")
  number:=0
  total:=0
  for number!=-1{
    total+=number
    fmt.Scan(&number)
  }
  fmt.Println(total)

  //Foreach Loop (You need to learn array first)
  fmt.Println("Foreach Loop")
  arr := []int{5,6,7,4,3,3,2,1,2,3,4,5,6}
  for _,k := range(arr){
    fmt.Println(k)
  }


}
