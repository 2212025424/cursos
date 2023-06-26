package handler

import (
	"errors"

	"github.com/2212025424/api/model"
)

type StorageMockError struct{}

func (sme *StorageMockError) Create(person *model.Person) error {
	return errors.New("error de mock")
}
func (sme *StorageMockError) Update(ID int, person *model.Person) error {
	return errors.New("error de mock")
}
func (sme *StorageMockError) Delete(ID int) error {
	return errors.New("error de mock")
}
func (sme *StorageMockError) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("error de mock")
}
func (sme *StorageMockError) GetAll() (model.Persons, error) {
	return model.Persons{}, errors.New("error de mock")
}

type StorageMockOk struct{}

func (sme *StorageMockOk) Create(person *model.Person) error {
	return nil
}
func (sme *StorageMockOk) Update(ID int, person *model.Person) error {
	return nil
}
func (sme *StorageMockOk) Delete(ID int) error {
	return nil
}
func (sme *StorageMockOk) GetByID(ID int) (model.Person, error) {
	return model.Person{}, nil
}
func (sme *StorageMockOk) GetAll() (model.Persons, error) {
	return model.Persons{}, nil
}
