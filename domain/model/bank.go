package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" valid:"notnull"`
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func NewBank(code, name string) (*Bank, error) {

	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewString()
	bank.CreatedAt = time.Now()

	bank.isValid()

	return &bank, nil
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}
