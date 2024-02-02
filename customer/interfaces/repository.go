package customer

import "CleanArch/customer"

type customerRepository struct {
	customer customer.CustomerModel
}

func NewCustomerRepository(Cus customer.CustomerModel) CustomerRepository {
	return &customerRepository{customer: Cus}
}

func (c *customerRepository) GetCustomerInfo() customer.CustomerModel {
	return c.customer
}
