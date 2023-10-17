package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vttrz/codepix/domain/model"
)

func TestNewAccount(t *testing.T) {

	bank, _ := model.NewBank("001", "Fake Bank")
	number := "1"
	ownerName := "fake user"

	t.Run("should create a account successfully", func(t *testing.T) {

		account, err := model.NewAccount(bank, number, ownerName)

		assert.Nil(t, err)
		assert.NotNil(t, account)
		assert.Equal(t, number, account.Number)
		assert.Equal(t, ownerName, account.OwnerName)

	})

	t.Run("should return an error because of invalid number", func(t *testing.T) {

		account, err := model.NewAccount(bank, "", ownerName)

		assert.Nil(t, account)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "number of account cannot be null")

	})
}
