package customer

import (
	. "CleanArch/customer/database"
	"github.com/kataras/iris/v12"
)

type Repository interface {
	GetCustomerInfo() CustomerModel
}

type UseCase interface {
	GetCustomerInfo() CustomerModel
}

type Controller interface {
	GetCustomerInfo(ctx iris.Context)
}
