package main
import "fmt"

func main()  {
  var name string = "Asif"
  fmt.Println("Your name is",name+".")
  var age int=24
  fmt.Println("Your age is",age)
  myName := "Asif"
  fmt.Println("Your name is",myName+".")
  myAge :=24
  fmt.Println("Your age is",myAge)
  var pi float64=3.14159265
  fmt.Println("Value of PI=",pi)
  fmt.Printf("Value of PI is %.3f\n",pi)
  fmt.Printf("DataType of PI is %T\n",pi)
  fmt.Printf("Hex of age is %x\n",myAge)
}
