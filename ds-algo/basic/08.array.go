package main
import "fmt"

func main(){
	var num[5] float64
  num[0]=1
  num[1]=11
  num[2]=34
  num[3]=55
  num[4]=5.66
  fmt.Println(num)
  fmt.Println(num[3])
  num2:=[3]float64{1,2,3}
  fmt.Println(num2[2])
  for i,value:=range num2{
    fmt.Println(value,i)
  }
  slice := num[3:5]
  fmt.Println("slice[1]=",slice[1])
  fmt.Println("num[:2]=",num[:2])
  fmt.Println("num[2:]=",num[2:])
  slice2 := make([]float64, 5,10)
  copy(slice2,slice)
  fmt.Println("slice2[1]=",slice2[1])
  slice2=append(slice2,0,-1)
  fmt.Println("slice2[6]=",slice2[6])
}
