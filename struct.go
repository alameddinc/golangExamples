package main

import "fmt"

type home struct{
  phone string
  address string
  motherName string
  fatherName string
}

type student struct{
  name string
  studentId string
  midtermE int
  finalE int
  makeupE int
  homeData home //point to consider
}

func main()  {
  firstHome:=home{"5554443322","1055 sk. Turkey","Mother","Father"}
  firstStudent := student{"Alameddin Çelik","3541",86,90,0,firstHome} //firstHome added
  //firstStudent.makeupE=100
  fmt.Println(firstStudent.finalE)
  fmt.Println(firstStudent.homeData.fatherName)
}
