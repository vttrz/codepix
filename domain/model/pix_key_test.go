package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vttrz/codepix/domain/model"
)

func TestNewPixKey(t *testing.T) {

	t.Run("should create a successfully pix key", func(t *testing.T) {
		code := "001"
		name := "Fake Bank"
		bank, _ := model.NewBank(code, name)

		accountNumber := "123"
		ownerName := "Fake User"

		account, _ := model.NewAccount(bank, accountNumber, ownerName)

		kind := "email"
		key := "fakeuser@mail.com"

		pixKey, err := model.NewPixKey(kind, account, key)

		assert.Nil(t, err)
		assert.NotEmpty(t, pixKey.ID)
		assert.Equal(t, pixKey.Kind, kind)
		assert.Equal(t, pixKey.Status, "active")

		kind = "cpf"
		_, err = model.NewPixKey(kind, account, key)

		assert.Nil(t, err)

	})
}
