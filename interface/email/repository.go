package email

import "pitest/entity"

type OutputEmailRepository interface {
	WriteOut([]*entity.OutputEmail) error
}

type ErrorCustomerRepository interface {
	WriteOut([]*entity.Customer) error
}
