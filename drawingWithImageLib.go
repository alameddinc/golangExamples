package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
)
var width=100
var height=100
var img = image.NewRGBA(image.Rect(0, 0, width, height))
var col color.Color

func main() {
    col = color.RGBA{255, 0, 0, 255} // Red
    setBackground(color.RGBA{255,255,255,255})
    HLine(10, 20, 80)
    Rect(10,10,90,90)
    f, err := os.Create("draw.png")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    png.Encode(f, img)
}

func setBackground(c color.Color)  {
  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      img.Set(x,y,c)
    }
  }
}

func HLine(x1, y, x2 int) {
    for ; x1 <= x2; x1++ {
        img.Set(x1, y, col)
    }
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
    for ; y1 <= y2; y1++ {
        img.Set(x, y1, col)
    }
}

func addLabel(img *image.RGBA, x, y int, label string) {
    c.SetDst(img)
    size := 12.0 // font size in pixels
    pt := freetype.Pt(x, y+int(c.PointToFixed(size)>>6))

    if _, err := c.DrawString(label, pt); err != nil {
        // handle error
    }
}


// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
    HLine(x1, y1, x2)
    HLine(x1, y2, x2)
    VLine(x1, y1, y2)
    VLine(x2, y1, y2)
}
