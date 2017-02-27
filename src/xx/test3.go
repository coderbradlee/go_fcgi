package main

import (
	"fmt"
	"runtime"
)
type Optioner interface{
	Name() string
	IsValid() bool
}
type OptionCommon struct{
	ShortName string "short option name"
	LongName string "long option name"
}
type IntOption struct{
	OptionCommon
	Value,Min,Max int
}
func (i *IntOption)Name()string {
	return i.ShortName+":"+i.LongName
}
func (i *IntOption)IsValid()bool {
	return i.Min<i.Max
}
type Exchanger interface{
	Exchange()
}
type Place struct{
	latitude,longitude float64
	Name string
}
func (place *Place)Exchange() {
	place.latitude,place.longitude=place.longitude,place.latitude
}
func New(latitude,longitude float64,name string)(*Place) {
	return &Place{latitude,longitude,name}
}
func (place *Place)SetLatitude(latitude float64){
	place.latitude=latitude
}
func (place *Place)SetLongitude(longitude float64) {
	place.longitude=longitude
}
func (place *Place)GetLatitude()float64 {
	return place.latitude
}
func (place *Place)GetLongitude()float64 {
	return place.longitude
}
func (place *Place)String()string {
	return fmt.Sprintf("%.3f,%.3f %q",place.latitude,place.longitude,place.Name)
}

func main() {
	// t:=New(1,2,"xx")
	// fmt.Println(t)
	// t.Exchange()
	// fmt.Println(t)
	// i:=IntOption{OptionCommon:OptionCommon{"s","long"},Max:10}
	// switch option:=i.(type){
	// case IntOption:
	// 	fmt.Println(i.Name())
	// }
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.GOMAXPROCS())
}
