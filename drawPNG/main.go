package main

import (
  "math/rand"
  "time"
  "fmt"
  "github.com/fogleman/gg"
)


func main()  {
  r := rand.New(rand.NewSource(99))
  r.Seed(time.Now().UTC().UnixNano())

  drawItem:= gg.NewContext(1000,1000)


  drawItem.SetHexColor("#fff")
  drawItem.DrawRectangle(0,0,1000,1000)
  drawItem.Fill()
  for i := 0; i < 40; i++ {
    for j := 0; j < 40; j++ {
      randColor:=r.Intn(16000000)
      randSize:=r.Intn(25)
      randomHex:=fmt.Sprintf("#%.6X", randColor)
      fmt.Println(randomHex)
      drawItem.SetHexColor(randomHex)
      drawItem.DrawCircle(float64(i*30),float64(j*30),float64(randSize))
      drawItem.Fill()

    }
  }

  defer drawItem.SavePNG("out3.png")

}
