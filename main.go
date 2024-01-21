package main

import (
	"CleanArch/Customer/infastructure"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	infastructure.Router(app)
	err := app.Listen(":8081")
	if err != nil {
		return
	}
}
