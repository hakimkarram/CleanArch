package customer

import . "CleanArch/customer/database"

type repository struct {
	customer CustomerModel
}

func NewRepository(Cus CustomerModel) Repository {
	return &repository{customer: Cus}
}

func (c *repository) GetCustomerInfo() CustomerModel {
	return c.customer
}
