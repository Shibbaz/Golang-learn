package main

import "fmt"

func main(){
	rectangle := Rectangle{
		Width: 23.2,
		Height: 32,
	}

	fmt.Println(Area(rectangle))
}