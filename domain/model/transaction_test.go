package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vttrz/codepix/domain/model"
)

func TestNewTransaction(t *testing.T) {

	t.Run("should create a transaction successfully", func(t *testing.T) {
		code := "001"
		name := "Fake Bank"
		bank, _ := model.NewBank(code, name)

		accountNumber := "123"
		ownerName := "Fake user"
		account, _ := model.NewAccount(bank, accountNumber, ownerName)

		accountNumberDestination := "abcdestination"
		ownerName = "Fake user 2"
		accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

		kind := "email"
		key := "j@j.com"
		pixKey, _ := model.NewPixKey(kind, accountDestination, key)

		require.NotEqual(t, account.ID, accountDestination.ID)

		amount := 3.10
		statusTransaction := "pending"
		transaction, err := model.NewTransaction(account, amount, pixKey, "My description")

		require.Nil(t, err)
		require.NotNil(t, transaction.ID)
		require.Equal(t, transaction.Amount, amount)
		require.Equal(t, transaction.Status, statusTransaction)
		require.Equal(t, transaction.Description, "My description")
		require.Empty(t, transaction.CancelDescription)

	})

	t.Run("should change transaction status", func(t *testing.T) {
		code := "001"
		name := "Fake Bank"
		bank, _ := model.NewBank(code, name)

		accountNumber := "123"
		ownerName := "Fake user"
		account, _ := model.NewAccount(bank, accountNumber, ownerName)

		accountNumberDestination := "abcdestination"
		ownerName = "Fake user 2"
		accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

		kind := "email"
		key := "j@j.com"
		pixKey, _ := model.NewPixKey(kind, accountDestination, key)

		amount := 3.10
		transaction, _ := model.NewTransaction(account, amount, pixKey, "My description")

		transaction.Complete()

		assert.Equal(t, transaction.Status, model.TransactionCompleted)

		transaction.Error("Error")

		assert.Equal(t, transaction.Status, model.TransactionError)
		assert.Equal(t, transaction.CancelDescription, "Error")
	})

	t.Run("should not be able to send pix to same account", func(t *testing.T) {
		code := "001"
		name := "Fake Bank"

		bank, err := model.NewBank(code, name)
		assert.Nil(t, err)

		accountNumber := "123"
		ownerName := "Fake user"

		account, err := model.NewAccount(bank, accountNumber, ownerName)
		assert.Nil(t, err)

		kind := "email"
		key := "j@j.com"

		pixKey, err := model.NewPixKey(kind, account, key)
		assert.Nil(t, err)

		amount := 3.10

		transaction, err := model.NewTransaction(account, amount, pixKey, "My description")

		assert.NotNil(t, err)
		assert.Nil(t, transaction)
		assert.Equal(t, err.Error(), "the source and destination account cannot be the same")
	})
}
