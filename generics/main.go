package main

import (
	"fmt"
	. "libs"
)

func main() {
	fmt.Println("Generics\n")
	router := Router{
		BaseAddr: "xd",
	}
	router.SetBaseAddr("localhost")
	fmt.Println("new baseaddr ", router.GetBaseAddr())
	routerGateway := Gateway[*Router]{
		Data: &router,
	}
	Print(routerGateway.Reflect())

	fmt.Println("\nApp example!")
	app := App{
		Port: 8080,
	}
	app.SetPort(9000)
	fmt.Println("new port ", app.GetPort())
	appGateway := Gateway[*App]{
		Data: &app,
	}
	Print(appGateway.Reflect())
}

