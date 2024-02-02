package customer

import "gorm.io/gorm"

type CustomerModel struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	Age          uint   `json:"age"`
	Address      string `json:"address"`
}
