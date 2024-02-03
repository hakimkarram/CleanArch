package customer

import (
	"github.com/kataras/iris/v12"
)

type controller struct {
	userCase UseCase
}

func NewController(uc UseCase) Controller {
	return &controller{userCase: uc}
}

func (c controller) GetCustomerInfo(ctx iris.Context) {
	err := ctx.JSON(iris.Map{"data": c.userCase.GetCustomerInfo()})
	if err != nil {
		return
	}
}
