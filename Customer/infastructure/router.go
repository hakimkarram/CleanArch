package infastructure

import (
	"CleanArch/Customer/interfaces"
	"github.com/kataras/iris/v12"
)

type router struct {
	app   *iris.Application
	party iris.Party
}

type Router interface {
	launch()
}

func NewRouter(builder *iris.Application, partyName string) Router {
	customerParty := builder.Party("/" + partyName)
	customerRouter := &router{app: builder, party: customerParty}
	customerRouter.launch()
	return customerRouter
}

func (c *router) launch() {
	c.party.Get("", interfaces.GetCustomerInfoHandler)
	c.party.Get("/greeting", interfaces.GetCustomerGreetingHandler)
}
