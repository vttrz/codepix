package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vttrz/codepix/domain/model"
)

func TestNewBank(t *testing.T) {

	t.Run("should create a bank successfully", func(t *testing.T) {

		code := "001"
		name := "Fake Bank"

		bank, err := model.NewBank(code, name)

		assert.Equal(t, code, bank.Code)
		assert.Equal(t, name, bank.Name)
		assert.Nil(t, err)
	})
}
