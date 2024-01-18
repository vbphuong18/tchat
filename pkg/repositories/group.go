package repositories

import (
	"TChat/pkg/domain"
	"TChat/pkg/models"
	"gorm.io/gorm"
)

type group struct {
	db *gorm.DB
}
type GroupRepository interface {
	CreateGroup(group domain.Group) error
	ListFriend(userID string) ([]models.Friend, error)
	DeleteFriend(userID1 string, userID2 string) error
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &group{
		db: db,
	}
}
