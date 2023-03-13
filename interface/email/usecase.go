package email

import "pitest/entity"

type EmailUsecase interface {
	Compose(
		template *entity.EmailTemplate,
		customers []*entity.Customer,
	) ([]*entity.OutputEmail, []*entity.Customer, error)
}
