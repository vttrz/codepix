package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `json:"bank" valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func NewAccount(bank *Bank, number, ownerName string) (*Account, error) {

	account := Account{
		Bank:      bank,
		OwnerName: ownerName,
		Number:    number,
	}

	account.ID = uuid.NewString()
	account.CreatedAt = time.Now()

	account.isValid()

	return &account, nil
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}
