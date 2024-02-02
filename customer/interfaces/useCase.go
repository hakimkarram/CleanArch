package customer

import "CleanArch/customer"

type customerUseCase struct {
	customerRepo CustomerRepository
}

func NewCustomerUseCase(repo CustomerRepository) CustomerUseCase {
	return &customerUseCase{customerRepo: repo}
}

func (c customerUseCase) GetCustomerInfo() customer.CustomerModel {
	return c.customerRepo.GetCustomerInfo()
}
