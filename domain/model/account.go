package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `gorm:"column:owner_name;type:varchar(255):not null" json:"owner_name" valid:"notnull"`
	Bank      *Bank     `json:"bank" valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string    `gorm:"type:varchar(255)" json:"number" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"`
}

func NewAccount(bank *Bank, number, ownerName string) (*Account, error) {

	account := Account{
		Bank:      bank,
		OwnerName: ownerName,
		Number:    number,
	}

	account.ID = uuid.NewString()
	account.CreatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if account.Number == "" {
		return errors.New("number of account cannot be null")
	}

	if err != nil {
		return err
	}

	return nil
}
