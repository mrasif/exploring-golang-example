package main

import "fmt"

func main() {
    var name string
	fmt.Print("Your name ? ")
    fmt.Scanf("%s",&name)
    fmt.Println("Hello",name,"!")
}
