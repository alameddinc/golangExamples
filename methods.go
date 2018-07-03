package main

import "fmt"

type robot struct{
  codeName string
  status int
  workText string
}

func (r *robot) whatsUrName()  {
  fmt.Println("Terminal : Version")
  fmt.Println("Robot:",r.codeName)
}
func (r *robot) isBusy()  {
  fmt.Println("Terminal : Status")
  fmt.Println("Robot:",r.status==1)
}
func (r *robot) work(workText string)  {
  fmt.Println("Terminal : Work For",workText)
  if r.status==0{
    r.status=1
    r.workText=workText
    fmt.Println("Robot: I'm working for",workText)
  }else{
    fmt.Println("Robot: I'm sorry, You must stop me first")
  }
}
func (r *robot) stop()  {
  fmt.Println("Terminal : Stop!")
  if r.status==1{
    r.status=0
    r.workText=""
  }else{
    fmt.Println("Robot: I'm sorry, I dont have a job")
  }
}

func main()  {
  yeni := robot{"47x65",0,""}

  yeni.whatsUrName()
  yeni.work("cleaning")
  yeni.work("production")
  yeni.stop()
  yeni.work("production")
}
