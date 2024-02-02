package customer

import (
	"CleanArch/customer"
	"github.com/kataras/iris/v12"
)

type CustomerRepository interface {
	GetCustomerInfo() customer.CustomerModel
}

type CustomerUseCase interface {
	GetCustomerInfo() customer.CustomerModel
}

type CustomerController interface {
	GetCustomerInfo(ctx iris.Context)
}
