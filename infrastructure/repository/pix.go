package repository

import (
	"errors"

	"github.com/vttrz/codepix/domain/model"
	"gorm.io/gorm"
)

type PixKeyRepositoryDB struct {
	db *gorm.DB
}

func (r *PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.db.Create(pixKey).Error

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (r *PixKeyRepositoryDB) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.db.Preload("Account.Bank").First(&pixKey, "key = ? and kind = ?", key, kind)

	if pixKey.ID == "" {
		return nil, errors.New("no key was found")
	}

	return &pixKey, nil
}

func (r *PixKeyRepositoryDB) AddBank(bank *model.Bank) error {

	err := r.db.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PixKeyRepositoryDB) AddAccount(account *model.Account) error {

	err := r.db.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, errors.New("no account was found")
	}

	return &account, nil
}

// Ele fez um método: FindBank
