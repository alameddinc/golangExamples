package main

import "fmt"

func main()  {
  adminName:="Admin"
  adminPass:="1289"
  enteredName:="Admin"
  enteredPass:="1234"


  //if & else
  if adminName==enteredName && adminPass==enteredPass{
    fmt.Println("It's True")
  }else{
    fmt.Println("It's False")
  }

  //switch
  switch enteredName {
  case "admin": fmt.Println("Welcome my Lord")
  case "moderator": fmt.Println("Welcome -_-")
  default: fmt.Println("Who are you!!!")

  }
}
