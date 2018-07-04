package main

import (
  "fmt"
  "plateLib"
)


func main()  {
  plate := plateLib.Plate{}
  plate.InitPlate("Çiçek")

  fmt.Printf("%s\n",plate.GetName())

  fmt.Printf("Water: %.2f\n",plate.GetWater())
  plate.Water()
  fmt.Printf("Water: %.2f\n",plate.GetWater())

  fmt.Printf("Mineral :%.2f\n",plate.GetMineral())
  plate.Mineral()
  fmt.Printf("Mineral :%.2f\n",plate.GetMineral())
}
