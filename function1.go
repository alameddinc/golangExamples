package main

import (
  "fmt"
  "strings"
  )

func sayMeHi()  {
  fmt.Println("Hi")
}

func sumFunc(num1,num2 int){
  fmt.Println(num1+num2)
}

func sumFunc2(num1,num2 int) int{
  return num1+num2
}

func halfAndDouble (num1 int)(float64,int){
  return float64(num1)/2,num1*2
}


func sumAll(nums ...int) int {
  total:=0
  for _,num:=range(nums){
    total+=num
  }
  return total
}


func main()  {
  sayMeHi()
  sumFunc(2,5)

  i:=sumFunc2(5,2)
  fmt.Println(i)

  halfData,doubleData:=halfAndDouble(5)
  fmt.Println(halfData," ",doubleData)

  result:=sumAll(2,3,4,5,7,10)
  fmt.Println(result)
  cyrpterText("zalameddin",27)


}
