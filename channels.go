package main

import "fmt"
  func sumP(arr []int,c chan int )  {
    myTotal:=0
    for _,k:=range(arr){
      myTotal+=k
    }
    c<-(myTotal*myTotal)
  }

func main() {
  c:= make(chan int)
  arr := makeRange(0, 10000)

  go sumP(arr[:2*len(arr)/3],c)
  go sumP(arr[len(arr)/3:len(arr)/3],c)
  go sumP(arr[2*len(arr)/3:],c)

  x,y,z:=<-c,<-c,<-c
  fmt.Println(x+y+z)
}


func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}
