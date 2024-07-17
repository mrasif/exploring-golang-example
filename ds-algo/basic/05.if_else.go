package main

import "fmt"

func main(){
	yourAge:=18
	if yourAge >= 16{
		fmt.Println("You can Drive")
	} else if yourAge >= 18 {
		fmt.Println("You can Vote")
	} else {
		fmt.Println("You can have Fun")
	}
}
