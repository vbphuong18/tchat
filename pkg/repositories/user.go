package repositories

import (
	"TChat/pkg/domain"
	"TChat/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func (u *user) CreateUser(user domain.User) error {
	return u.db.Create(&models.User{
		UserID:      uuid.New().String(),
		PhoneNumber: user.PhoneNumber,
		DateOfBirth: user.DateOfBirth,
		Name:        user.Name,
		Email:       user.Email,
		Gender:      models.GenderType(user.Gender),
		UserName:    user.UserName,
		Password:    user.Password,
		AvtImg:      user.AvtImg,
		CoverImg:    user.CoverImg,
	}).Error
}

func (u *user) ListUser() ([]models.User, error) {
	var us []models.User
	err := u.db.Find(&us).Error
	return us, err
}

func (u *user) SearchUser(name string, phoneNumber string) ([]models.User, error) {
	var search []models.User
	err := u.db.Where("name like ? or phone_number = ?", "%"+name+"%", phoneNumber).Find(&search).Error
	return search, err
}

func (u *user) DeleteUser(userID string) error {
	return u.db.Where("user_id = ?", userID).Delete(&models.User{}).Error
}

type UserRepository interface {
	CreateUser(user domain.User) error
	ListUser() ([]models.User, error)
	SearchUser(name string, phoneNumber string) ([]models.User, error)
	DeleteUser(userID string) error
} // define method func

func NewUserRepository(db *gorm.DB) UserRepository {
	return &user{
		db: db,
	}
} // init object func
