package main

import (
	"log"
	"os"
	"pitest/impl/email"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("Invalid arguments")
	}
	// init repositories
	outputEmailRepo := email.NewOutputEmailJsonRepo(os.Args[3])
	errorCustomerCsvRepo := email.NewErrorCustomerCsvRepo(os.Args[4])

	// init usecase
	usecase := email.NewEmailUsecase(&outputEmailRepo, &errorCustomerCsvRepo)

	// init controller
	ffEmailHandler := email.NewFromFileEmailHandler(&usecase)

	if err := ffEmailHandler.ComposeEmail(); err != nil {
		log.Fatal(err)
	}
}
