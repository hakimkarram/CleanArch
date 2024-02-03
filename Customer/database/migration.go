package consumer

import (
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) (*gorm.DB, error) {
	err1 := db.AutoMigrate(&CustomerModel{})
	if err1 != nil {
		return nil, err1
	}
	return db, nil
}

func GetDBCustomInfo() CustomerModel {
	return CustomerModel{
		CustomerName: "Hakim Adil Kadhim",
		Age:          38,
		Address:      "Najaf - Ameer Quarter",
	}
}
