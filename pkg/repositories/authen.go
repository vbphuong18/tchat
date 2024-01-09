package repositories

import (
	"TChat/pkg/models"
	"gorm.io/gorm"
)

type authen struct {
	db *gorm.DB
}

func (a *authen) GetUserByUserName(userName string) (models.User, error) {
	var auth models.User
	err := a.db.Where("user_name = ?", userName).First(&auth).Error
	return auth, err
}

type AuthenRepository interface {
	GetUserByUserName(userName string) (models.User, error)
}

func NewAuthenRepository(db *gorm.DB) AuthenRepository {
	return &authen{
		db: db,
	}
}
