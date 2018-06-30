package main

import "fmt"

func main()  {
  //Method 1
  var arr [3]float64

  arr[0]=3.54
  arr[1]=-4.33
  arr[2]=14.53

  //Method 2 (Slice Functions)
  var arr2 []int
  arr2 = append(arr2,5)
  arr2 = append(arr2,3)
  arr2 = append(arr2,14)

  var arr21 []int= append(arr2,3,4)

  //Method 3
  arr3:= [12]float64{1,2,3,4,5}

  //Method 4
  arr4:= []float64{1,2,3,4,5,6,7,8,9,122}

  fmt.Println(arr)
  fmt.Println(arr2)
  fmt.Println(arr21)
  fmt.Println(arr3)
  fmt.Println(arr4)


}
