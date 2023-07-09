package main

import (
	"fmt"
	"reflect"
)

type App struct {
	Number int
}

func (app *App) Reflect()(reflect.Type, reflect.Type){
	return reflect.TypeOf(app.Number), reflect.TypeOf(app)
}

type Router struct {
	Number string
}

type ExtendedRouter struct {
	App reflect.Type
	Value reflect.Type
}

func (router *Router) Reflect()(reflect.Type, reflect.Type){
	ExtendedRouter := ExtendedRouter{
		Value: reflect.TypeOf(router.Number),
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

func (text Gateway[Source]) Reflect()(reflect.Type, reflect.Type)  {
	return text.Data.Reflect()
}

func main() {
	fmt.Println("Generic\n")

	fmt.Println("\nRouter example!")
	router := Router{
		Number: "xd",
	}
	routerGateway := Gateway[*Router]{
		Data: &router,
	}
	routerTypeA, routerTypeB := routerGateway.Reflect()
	fmt.Println(routerTypeA)
	fmt.Println(routerTypeB)

	fmt.Println("\nApp example!")
	app := App{
		Number: 1,
	}
	appGateway := Gateway[*App]{
		Data: &app,
	}
	appTypeA, appTypeB := appGateway.Reflect()
	fmt.Println(appTypeA)
	fmt.Println(appTypeB)
}

