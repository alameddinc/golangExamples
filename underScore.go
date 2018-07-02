package main

import "fmt"

func getExamsResult(studentId int) (int,int,int)  {
  midtermE:=90
  finalE:=24
  makeupE:=100
  return midtermE,finalE,makeupE
}

func main()  {
  //If I want to fetch only finalE
  _,f,_:=getExamsResult(12323123)

  fmt.Println("Student's final grade: ",f)
}
