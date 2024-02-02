package settings

import (
	"CleanArch/customer/infastructure"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func Settings(App *iris.Application, db *gorm.DB) {
	api := App.Party("/api")
	customer.LaunchApp("customer", db, api)
}
