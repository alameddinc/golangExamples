package main

import "fmt"

type robot struct{
  codeName string
  status int
}

func (r *robot) whatsUrName()  {
  fmt.Println(r.codeName)
}
func (r *robot) isBusy()  {
  fmt.Println(r.status==1)
}
func (r *robot) work()  {
  if r.status==0{
    r.status=1
  }else{
    fmt.Println("I'm sorry, You must stop me first")
  }
}
func (r *robot) stop()  {
  if r.status==1{
    r.status=0
  }else{
    fmt.Println("I'm sorry, I dont have a job")
  }
}

func main()  {
  yeni := robot{"47x65",0}
  yeni.isBusy()
  yeni.work()
  yeni.work()
  yeni.isBusy()
}
