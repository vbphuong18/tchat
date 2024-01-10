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

type FriendRepository interface {
	CreateFriend(friend domain.Friend) error
	//ListFriend
	//DeleteFriend

}

func NewFriendRepository(db *gorm.DB) FriendRepository {
	return &friend{
		db: db,
	}
}
