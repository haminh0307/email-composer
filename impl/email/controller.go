package email

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"pitest/entity"
	"pitest/interface/email"
)

type FromFileEmailHandler struct {
	emailUC email.EmailUsecase
}

func NewFromFileEmailHandler(emailUC email.EmailUsecase) FromFileEmailHandler {
	return FromFileEmailHandler{
		emailUC: emailUC,
	}
}

func LoadTemplateFromJsonFile(templateJsonPath string) (*entity.EmailTemplate, error) {
	templateFile, err := os.Open(templateJsonPath)
	if err != nil {
		return nil, err
	}

	var template entity.EmailTemplate
	jsonParser := json.NewDecoder(templateFile)
	if err = jsonParser.Decode(&template); err != nil {
		return nil, err
	}

	return &template, err
}

func LoadCustomersFromCsvFile(customersCsvPath string) ([]*entity.Customer, error) {
	f, err := os.Open(customersCsvPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var customers []*entity.Customer
	for _, record := range records {
		// ignore if record does not contains exactly 4 fields
		if len(record) != 4 {
			continue
		}
		customer := entity.Customer{
			Title:     record[0],
			FirstName: record[1],
			LastName:  record[2],
			Email:     record[3],
		}
		customers = append(customers, &customer)
	}

	return customers, nil
}

func (ffEmailHandler *FromFileEmailHandler) ComposeEmail() error {
	// load template from json file
	templateJsonPath := os.Args[1]
	template, err := LoadTemplateFromJsonFile(templateJsonPath)
	if err != nil {
		return err
	}

	// load customers from csv file
	customersCsvPath := os.Args[2]
	customers, err := LoadCustomersFromCsvFile(customersCsvPath)
	if err != nil {
		return err
	}

	_, _, err = ffEmailHandler.emailUC.Compose(template, customers)

	return err
}
