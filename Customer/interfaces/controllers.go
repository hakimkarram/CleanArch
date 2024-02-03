package customer

import (
	"github.com/kataras/iris/v12"
)

type customerController struct {
	customerUserCase CustomerUseCase
}

func NewCustomerController(CustomerUseCase CustomerUseCase) CustomerController {
	return &customerController{customerUserCase: CustomerUseCase}
}

func (c customerController) GetCustomerInfo(ctx iris.Context) {
	err := ctx.JSON(iris.Map{"data": c.customerUserCase.GetCustomerInfo()})
	if err != nil {
		return
	}
}
