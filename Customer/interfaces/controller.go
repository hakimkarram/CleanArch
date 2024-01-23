package interfaces

import (
	"CleanArch/Customer"
	"github.com/kataras/iris/v12"
)

func GetCustomerInfoHandler(ctx iris.Context) {
	err := ctx.JSON(iris.Map{"data": Customer.GetCustomInfo()})
	if err != nil {
		return
	}
}

func GetCustomerGreetingHandler(ctx iris.Context) {
	err := ctx.JSON(iris.Map{"data": "Hi from Hakim"})
	if err != nil {
		return
	}
}
