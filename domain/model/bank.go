package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name     string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

func NewBank(code, name string) (*Bank, error) {

	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewString()
	bank.CreatedAt = time.Now()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}
