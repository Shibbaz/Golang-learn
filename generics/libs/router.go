package libs

import (
	"reflect"
)

type ExtendedRouter struct {
	App reflect.Type
	Addr reflect.Type
}

type Router struct {
	BaseAddr string
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