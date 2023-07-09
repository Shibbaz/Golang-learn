package libs

import (
	"fmt"
	"reflect"
)

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