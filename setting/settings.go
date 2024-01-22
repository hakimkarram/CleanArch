package settings

import (
	Customer "CleanArch/Customer/infastructure"
	"github.com/kataras/iris/v12"
)

type settings struct {
	app *iris.Application
}

type Settings interface {
}

func NewSettings(App *iris.Application) Settings {
	Customer.NewRouter(App, "customer")
	return &settings{app: App}
}
