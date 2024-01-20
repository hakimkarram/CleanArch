package Customer

import (
	"github.com/kataras/iris/v12"
)

func GetCustomerInfoHandler(ctx iris.Context) {
	err := ctx.JSON(iris.Map{"data": GetCustomInfo()})
	if err != nil {
		return
	}
}
