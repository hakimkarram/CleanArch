package customer

import (
	. "CleanArch/customer/interface"
	"github.com/kataras/iris/v12"
)

type router struct {
	api iris.Party
}

type Router interface {
	launch(Controller)
}

func LaunchRouter(api iris.Party, AppName string, controller Controller) Router {
	apiPrefix := api.Party("/" + AppName)
	appRouter := &router{api: apiPrefix}
	appRouter.launch(controller)
	return appRouter
}

func (c *router) launch(controller Controller) {
	c.api.Get("", controller.GetCustomerInfo)
	c.api.Get("/greeting", controller.GetCustomerInfo)
}
