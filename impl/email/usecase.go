package email

import (
	"net/mail"
	"pitest/entity"
	"pitest/interface/email"
	"strings"
	"time"
)

const Layout = "02 Jan 2006"

type EmailUsecaseImpl struct {
	outputEmailRepo   email.OutputEmailRepository
	errorCustomerRepo email.ErrorCustomerRepository
}

func NewEmailUsecase(oEmailRepo email.OutputEmailRepository, errCustomerRepo email.ErrorCustomerRepository) EmailUsecaseImpl {
	return EmailUsecaseImpl{
		outputEmailRepo:   oEmailRepo,
		errorCustomerRepo: errCustomerRepo,
	}
}

func ReplacePlaceholders(src string, customer *entity.Customer) string {
	// replace with customer's information
	result := strings.ReplaceAll(src, "{{TITLE}}", customer.Title)
	result = strings.ReplaceAll(result, "{{FIRST_NAME}}", customer.FirstName)
	result = strings.ReplaceAll(result, "{{LAST_NAME}}", customer.LastName)

	// replace with current date
	result = strings.ReplaceAll(result, "{{TODAY}}", time.Now().Format(Layout))

	return result
}

func (emailUC *EmailUsecaseImpl) ComposeSingleOutputEmail(template *entity.EmailTemplate, customer *entity.Customer) (*entity.OutputEmail, error) {
	// check valid mail
	if _, err := mail.ParseAddress(customer.Email); err != nil {
		return nil, err
	}

	result := entity.OutputEmail{
		From:     template.From,
		To:       customer.Email,
		Subject:  template.Subject,
		MimeType: template.MimeType,
		Body:     ReplacePlaceholders(template.Body, customer),
	}

	return &result, nil
}

func (emailUC *EmailUsecaseImpl) Compose(template *entity.EmailTemplate, customers []*entity.Customer) ([]*entity.OutputEmail, []*entity.Customer, error) {
	var results []*entity.OutputEmail
	var errCustomers []*entity.Customer

	for _, customer := range customers {
		if oEmail, err := emailUC.ComposeSingleOutputEmail(template, customer); err == nil {
			results = append(results, oEmail)
		} else {
			errCustomers = append(errCustomers, customer)
		}
	}

	// write output emails
	if err := emailUC.outputEmailRepo.WriteOut(results); err != nil {
		return results, errCustomers, err
	}

	// write error customers
	if err := emailUC.errorCustomerRepo.WriteOut(errCustomers); err != nil {
		return results, errCustomers, err
	}

	return results, errCustomers, nil
}
