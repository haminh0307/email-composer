package email

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"pitest/entity"
)

type OutputEmailJsonRepo struct {
	pathToOutputEmails string
}

func NewOutputEmailJsonRepo(pathToOutputEmails string) OutputEmailJsonRepo {
	return OutputEmailJsonRepo{
		pathToOutputEmails: pathToOutputEmails,
	}
}

func (jsonRepo *OutputEmailJsonRepo) WriteOut(oEmails []*entity.OutputEmail) error {
	data, err := json.Marshal(oEmails)
	if err != nil {
		return err
	}

	outputPath := filepath.Join(jsonRepo.pathToOutputEmails, "output.json")
	err = ioutil.WriteFile(outputPath, data, os.ModeAppend)

	return err
}
