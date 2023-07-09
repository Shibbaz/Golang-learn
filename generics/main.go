package main

import (
	"fmt"
	"reflect"
)

type App struct {
	Port int
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
	routerGateway := Gateway[*Router]{
		Data: &router,
	}
	Print(routerGateway.Reflect())

	fmt.Println("\nApp example!")
	app := App{
		Port: 8080,
	}
	appGateway := Gateway[*App]{
		Data: &app,
	}
	Print(appGateway.Reflect())
}

