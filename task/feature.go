package task

import (
	"encoding/json"
	"os"
)

type Database struct {
	FilePath string
}

func (db Database) Add(person Person) error {
	data, err := db.List()
	if err != nil {
		return err
	}

	data = append(data, person)
	file, _ := os.OpenFile(db.FilePath, os.O_CREATE, os.ModeDevice)

	defer file.Close()
	err = json.NewEncoder(file).Encode(data)

	return err
}

func (db Database) List() ([]Person, error) {
	var path = db.FilePath
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var data []Person
	err = json.NewDecoder(file).Decode(&data)
	return data, err
}
