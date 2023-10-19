package usecase

import (
	"fmt"

	"github.com/vttrz/codepix/domain/model"
)

type TransactionUseCase struct {
	transactionRepo model.TransactionRepositoryInterface
	pixRepo         model.PixKeyRepositoryInterface
}

func (uc *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {

	account, err := uc.pixRepo.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := uc.pixRepo.FindKeyByKind(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	err = uc.transactionRepo.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil

}

func (uc *TransactionUseCase) findTransactionByID(transactionId string) (*model.Transaction, error) {
	transaction, err := uc.transactionRepo.Find(transactionId)

	if err != nil {
		return nil, fmt.Errorf("transaction was not found: %s", err)
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {

	transaction, _ := uc.findTransactionByID(transactionId)

	err := transaction.Confirm()

	if err != nil {
		return nil, err
	}

	err = uc.transactionRepo.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {

	transaction, _ := uc.findTransactionByID(transactionId)

	err := transaction.Complete()

	if err != nil {
		return nil, err
	}

	err = uc.transactionRepo.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Error(transactionId, reason string) (*model.Transaction, error) {

	transaction, _ := uc.findTransactionByID(transactionId)

	err := transaction.Error(reason)

	if err != nil {
		return nil, err
	}

	err = uc.transactionRepo.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
