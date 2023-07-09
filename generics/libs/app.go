package libs

import (
	"reflect"
)

type App struct {
	Port int
}

func NewAppGateway(app App) Gateway[*App]{
	return Gateway[*App]{
		Data: &app,
	}
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