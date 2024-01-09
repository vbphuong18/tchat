package services

import (
	"TChat/pkg/repositories"
)

type authen struct {
	authenRepository repositories.AuthenRepository
}

func (a *authen) Login(userName string) (string, error) {
	userModels, err := a.authenRepository.GetUserByUserName(userName)
	if err != nil {
		return "", err
	}
	return userModels.Password, nil
}

type AuthenService interface {
	Login(userName string) (string, error)
}

func NewAuthenService(authenRepository repositories.AuthenRepository) AuthenService {
	return &authen{
		authenRepository: authenRepository,
	}
}
