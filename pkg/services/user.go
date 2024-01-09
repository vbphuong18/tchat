package services

import (
	"TChat/pkg/domain"
	"TChat/pkg/repositories"
)

type user struct {
	userRepository repositories.UserRepository
}

func (u *user) CreateUser(user domain.User) error {
	return u.userRepository.CreateUser(user)
}

func (u *user) ListUser() ([]domain.User, error) {
	us, err := u.userRepository.ListUser()
	if err != nil {
		return nil, err
	}
	var usr []domain.User
	for i := 0; i < len(us); i++ {
		usr = append(usr, domain.User{
			UserID:      us[i].UserID,
			DateOfBirth: us[i].DateOfBirth,
			Name:        us[i].Name,
			Gender:      domain.GenderType(us[i].Gender),
			AvtImg:      us[i].AvtImg,
			CoverImg:    us[i].CoverImg,
		})
	} // convert us(models) to usr(domain)
	return usr, nil
}

func (u *user) SearchUser(name string, phoneNumber string) ([]domain.User, error) {
	search, err := u.userRepository.SearchUser(name, phoneNumber)
	if err != nil {
		return nil, err
	}
	var srch []domain.User
	for i := 0; i < len(search); i++ {
		srch = append(srch, domain.User{
			UserID:      search[i].UserID,
			DateOfBirth: search[i].DateOfBirth,
			Name:        search[i].Name,
			Gender:      domain.GenderType(search[i].Gender),
			AvtImg:      search[i].AvtImg,
			CoverImg:    search[i].CoverImg,
		})
	} // convert search(models) to srch(domain)
	return srch, nil
}

func (u *user) DeleteUser(userID string) error {
	return u.userRepository.DeleteUser(userID)
}

type UserService interface {
	CreateUser(user domain.User) error
	ListUser() ([]domain.User, error)
	SearchUser(name string, phoneNumber string) ([]domain.User, error)
	DeleteUser(userID string) error
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &user{
		userRepository: userRepository,
	}
}
