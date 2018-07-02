package main

import (
  "fmt"
  "strings"
  )
  //Caesar Crypto
func cryptText (text string,code int) string{
  encodedText := text
  for _,k:=range text{
    encodedText = strings.Replace(encodedText, string(k), string(int(k)-code), -1)
  }
  return encodedText
}
func decryptText (text string,code int) string{
  decodedText := text
  for _,k:=range text{
    decodedText = strings.Replace(decodedText, string(k), string(int(k)+code), -1)
  }
  return decodedText
}


func main()  {
  fmt.Println("Text: ","Hello Mars")
  fmt.Println("Encoded: ",cryptText("Hello Mars",5))
  fmt.Println("Decoded: ",decryptText(cryptText("Hello Mars",5),5))


}
