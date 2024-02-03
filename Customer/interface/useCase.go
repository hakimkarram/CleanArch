package customer

import . "CleanArch/customer/database"

type useCase struct {
	repo Repository
}

func NewUseCase(repos Repository) UseCase {
	return &useCase{repo: repos}
}

func (c useCase) GetCustomerInfo() CustomerModel {
	return c.repo.GetCustomerInfo()
}
