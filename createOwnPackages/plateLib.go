//You need to add this file in $GOPATH
package plateLib

type Plate struct {
  Pid string
  wRatio float64 //water Ratio
  mRatio float64 //mineral Ratio
  oxygen float64
}

func (p *Plate) InitPlate(PId string) {
    p.Pid=PId
    p.wRatio=0.76
    p.mRatio=0.65
    p.oxygen=0.44532
}

func (p *Plate) Water() float64{
  for p.wRatio<0.8{
    //Open the water line (cycle)
    p.wRatio+=0.01;
  }
  return p.wRatio
}

func (p *Plate) Mineral()float64{
  for p.mRatio<0.8{
    //Open the mineral line (cycle)
    p.mRatio+=0.01;
  }
  return p.mRatio
}
func (p *Plate) oxySensor()  {
  p.oxygen = 0.76 //Analog to Digital
}
func (p *Plate) GetName() string {
  return p.Pid
}
func (p *Plate) GetMineral() float64 {
  return p.mRatio
}
func (p *Plate) GetWater() float64 {
  return p.wRatio
}
