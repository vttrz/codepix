package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {

	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewString()
	pixKey.CreatedAt = time.Now()

	pixKey.isValid()

	return &pixKey, nil
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "cpf" && pixKey.Kind != "email" {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}
