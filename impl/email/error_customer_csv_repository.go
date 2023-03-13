package email

import (
	"encoding/csv"
	"os"
	"pitest/entity"
)

type ErrorCustomerCsvRepo struct {
	pathToOutputError string
}

func NewErrorCustomerCsvRepo(pathToOutputError string) ErrorCustomerCsvRepo {
	return ErrorCustomerCsvRepo{
		pathToOutputError: pathToOutputError,
	}
}

func (csvRepo *ErrorCustomerCsvRepo) WriteOut(customers []*entity.Customer) error {
	f, err := os.OpenFile(csvRepo.pathToOutputError, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// write all the error customers
	for _, customer := range customers {
		if err = w.Write(customer.ToSliceOfStrings()); err != nil {
			return err
		}
	}

	return err
}
