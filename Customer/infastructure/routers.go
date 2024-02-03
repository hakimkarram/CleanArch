package customer

import (
	. "CleanArch/customer/interfaces"
	"github.com/kataras/iris/v12"
)

type router struct {
	api iris.Party
}

type Router interface {
	launch(CustomerController)
}

func LaunchRouter(api iris.Party, AppName string, controller CustomerController) Router {
	apiPrefix := api.Party("/" + AppName)
	appRouter := &router{api: apiPrefix}
	appRouter.launch(controller)
	return appRouter
}

func (c *router) launch(controller CustomerController) {
	c.api.Get("", controller.GetCustomerInfo)
	c.api.Get("/greeting", controller.GetCustomerInfo)
}
