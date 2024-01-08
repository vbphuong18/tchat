package services

import (
	"TChat/repositories"
	"TChat/utils"
)

type authen struct {
	authenRepository repositories.AuthenRepository
}

func (a *authen) Login(userName string, password string) error {
	user, err := a.authenRepository.GetUserByUserName(userName)
	if err != nil {
		return err
	}
	return utils.CheckPasswordHash(user.Password, password)
}

type AuthenService interface {
	Login(userName string, password string) error
}

func NewAuthenService(authenRepository repositories.AuthenRepository) AuthenService {
	return &authen{
		authenRepository: authenRepository,
	}
}
