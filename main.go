package main

import (
	"CleanArch/Customer"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	Customer.Router(app)

	err := app.Listen(":8081")

	if err != nil {
		return
	}
}
