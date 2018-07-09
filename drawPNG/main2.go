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
  drawItem.DrawCircle(500,500,400)
  drawItem.Fill()
  drawItem.DrawCircle(500,500,400)
  drawItem.SetRGBA(0, 0, 0, 0.2)
  drawItem.SetLineWidth(100)
  drawItem.Stroke()


  for i := 0; i <= 500; i++ {
    drawItem.SetHexColor(fmt.Sprintf("#%.3X",3840+r.Intn(100)))
    drawItem.RotateAbout(gg.Radians(float64(i*3)), 500, 500)
    drawItem.DrawLine(float64(i),float64(i),float64(i+5),float64(i+20))
    drawItem.SetLineWidth(r.Float64()*3)
    drawItem.Stroke()
  }

  drawItem.SetHexColor("#333")
  drawItem.DrawCircle(500,500,100)
  drawItem.Fill()

  defer drawItem.SavePNG("out3.png")

}
