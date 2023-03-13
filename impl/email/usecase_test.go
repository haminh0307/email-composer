package email_test

import (
	"pitest/entity"
	"pitest/impl/email"
	mock_email "pitest/interface/email/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCompose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mOutputEmailRepo := mock_email.NewMockOutputEmailRepository(ctrl)
	mErrorCustomerRepo := mock_email.NewMockErrorCustomerRepository(ctrl)

	UC := email.NewEmailUsecase(mOutputEmailRepo, mErrorCustomerRepo)

	template := entity.EmailTemplate{
		From:     "The Marketing Team<marketing@example.com>",
		Subject:  "A new product is being launched soon...",
		MimeType: "text/plain",
		Body:     "Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}}, \nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
	}
	customers := []*entity.Customer{
		{
			Title:     "Mr",
			FirstName: "John",
			LastName:  "Smith",
			Email:     "john.smith@example.com",
		},
		{
			Title:     "Mrs",
			FirstName: "Michelle",
			LastName:  "Smith",
			Email:     "michelle.smith@example.com",
		},
		{
			Title:     "Mr",
			FirstName: "Adam",
			LastName:  "Smith",
			Email:     "",
		},
	}

	expectedOutput := []*entity.OutputEmail{
		{
			From:     "The Marketing Team<marketing@example.com>",
			To:       "john.smith@example.com",
			Subject:  "A new product is being launched soon...",
			MimeType: "text/plain",
			Body:     "Hi Mr John Smith, \nToday, " + time.Now().Format(email.Layout) + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
		{
			From:     "The Marketing Team<marketing@example.com>",
			To:       "michelle.smith@example.com",
			Subject:  "A new product is being launched soon...",
			MimeType: "text/plain",
			Body:     "Hi Mrs Michelle Smith, \nToday, " + time.Now().Format(email.Layout) + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
		},
	}

	expectedErrorCustomers := []*entity.Customer{
		customers[2],
	}

	mOutputEmailRepo.EXPECT().WriteOut(expectedOutput).Return(nil)
	mErrorCustomerRepo.EXPECT().WriteOut(expectedErrorCustomers).Return(nil)

	actualOutput, actualErrorCustomers, err := UC.Compose(&template, customers)

	assert.Nil(t, err)
	assert.Equal(t, actualOutput, expectedOutput)
	assert.Equal(t, actualErrorCustomers, expectedErrorCustomers)
}
