package usecase

import (
	"fmt"

	"github.com/vttrz/codepix/domain/model"
)

type PixUseCase struct {
	repository model.PixKeyRepositoryInterface
}

func (uc *PixUseCase) RegisterKey(key, kind, accountId string) (*model.PixKey, error) {

	account, err := uc.repository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	_, err = uc.repository.RegisterKey(pixKey)

	if err != nil {
		return nil, fmt.Errorf("unable to create a new key at the moment: %s", err)
	}

	return pixKey, nil
}

func (uc *PixUseCase) FindKey(key, kind string) (*model.PixKey, error) {

	pixKey, err := uc.repository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
