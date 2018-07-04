package main

import "fmt"
  func sum(arr []int,c chan int )  {
    myTotal:=0
    for _,k:=range(arr){
      myTotal+=k
    }
    c<-myTotal
  }

func main() {
  c:= make(chan int)
  arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19}

  go sum(arr[:len(arr)/6],c)
  go sum(arr[len(arr)/3:len(arr)/3],c)
  go sum(arr[len(arr)/6:],c)

  x,y,z:=<-c,<-c,<-c
  fmt.Println(x+y+z)
}
