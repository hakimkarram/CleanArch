package infastructure

import (
	"CleanArch/Customer/interfaces"
	"github.com/kataras/iris/v12"
)

func Router(builder *iris.Application) {
	customer := builder.Party("/customer")
	customer.Get("", interfaces.GetCustomerInfoHandler)
}
