package customer

import (
	appDB "CleanArch/customer"
	. "CleanArch/customer/interfaces"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func LaunchApp(AppName string, db *gorm.DB, api iris.Party) {
	_, _ = appDB.MigrateDB(db)
	CostRepo := NewCustomerRepository(appDB.GetDBCustomInfo())
	CostUseCase := NewCustomerUseCase(CostRepo)
	CostController := NewCustomerController(CostUseCase)
	LaunchRouter(api, AppName, CostController)
}
