package main

import (
	"fmt"
	"reflect"
)

type App struct {
	Port int
}

func (app *App) SetPort(port int){
	app.Port = port
}

func (app *App) GetPort() int{
	return app.Port
}

func (app *App) Reflect()(reflect.Type, reflect.Type){
	return reflect.TypeOf(app.Port), reflect.TypeOf(app)
}

type Router struct {
	BaseAddr string
}

type ExtendedRouter struct {
	App reflect.Type
	Addr reflect.Type
}

func (router *Router) Reflect()(reflect.Type, reflect.Type){
	ExtendedRouter := ExtendedRouter{
		Addr: reflect.TypeOf(router.BaseAddr),
	}
	return reflect.TypeOf(ExtendedRouter), reflect.TypeOf(router)
}

func (router *Router) SetBaseAddr(baseAddr string){
	router.BaseAddr = baseAddr;
}

func (router *Router) GetBaseAddr() string{
	return router.BaseAddr
}

type Source interface {
	*App | *Router
	Reflect()(reflect.Type, reflect.Type)
}

type Gateway[T Source] struct {
    Data T
}

func (gateway Gateway[Source]) Reflect()(reflect.Type, reflect.Type)  {
	return gateway.Data.Reflect()
}

func Print(A reflect.Type, B reflect.Type){
	fmt.Println(A, B)
}

func NewAppGateway(app App) Gateway[*App]{
	return Gateway[*App]{
		Data: &app,
	}
}

func NewRouterGateway(router Router) Gateway[*Router]{
	return Gateway[*Router]{
		Data: &router,
	}
}

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

