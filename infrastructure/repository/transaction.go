package repository

import (
	"errors"

	"github.com/vttrz/codepix/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDB struct {
	db *gorm.DB
}

func (r *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	err := r.db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	err := r.db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, errors.New("no transaction was found")
	}

	return &transaction, nil
}
