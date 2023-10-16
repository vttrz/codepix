package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transactions []Transaction
}

type Transaction struct {
	Base              `valid:"-"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {

	transation := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}

	transation.ID = uuid.NewString()
	transation.CreatedAt = time.Now()

	transation.isValid()

	return &transation, nil
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("invalid status for the transaction")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()

	err := t.isValid()

	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()

	err := t.isValid()

	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.Description = description
	t.UpdatedAt = time.Now()

	err := t.isValid()

	if err != nil {
		return err
	}

	return nil
}
