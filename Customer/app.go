package customer

import (
	appDB "CleanArch/customer/database"
	router "CleanArch/customer/infastructure"
	. "CleanArch/customer/interface"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func LaunchApp(AppName string, db *gorm.DB, api iris.Party) {
	_, _ = appDB.MigrateDB(db)
	CostRepo := NewRepository(appDB.GetDBCustomInfo())
	CostUseCase := NewUseCase(CostRepo)
	CostController := NewController(CostUseCase)
	router.LaunchRouter(api, AppName, CostController)
}
