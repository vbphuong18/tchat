package repositories

import (
	"TChat/pkg/domain"
	"TChat/pkg/models"
	"gorm.io/gorm"
)

type friend struct {
	db *gorm.DB
}

func (f *friend) CreateFriend(fr domain.Friend) error {
	return f.db.Create(&models.Friend{
		UserID1: fr.UserID1,
		UserID2: fr.UserID2,
	}).Error
}
func (f *friend) ListFriend(userID string) ([]models.Friend, error) {
	var friendModel []models.Friend
	err := f.db.Where("user_id_1 = ? or user_id_2 = ?", userID, userID).Find(&friendModel).Error
	return friendModel, err
}

func (f *friend) DeleteFriend(userID1 string, userID2 string) error {
	return f.db.
		Where("user_id_1 = ? and user_id_2 = ?", userID1, userID2).Or("user_id_1 = ? and user_id_2 = ?",
		userID2, userID1).
		Delete(&models.Friend{}).Error
}

type FriendRepository interface {
	CreateFriend(friend domain.Friend) error
	ListFriend(userID string) ([]models.Friend, error)
	DeleteFriend(userID1 string, userID2 string) error
}

func NewFriendRepository(db *gorm.DB) FriendRepository {
	return &friend{
		db: db,
	}
}
