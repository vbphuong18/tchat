package repositories

import (
	"TChat/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type group struct {
	db *gorm.DB
}

func (g *group) CreateGroup(listUserID []string, name string, avt string) (models.Group, error) {
	tx := g.db.Begin()
	groupID := uuid.New().String()
	gr := models.Group{
		GroupID: groupID,
		Name:    name,
		Avt:     avt,
	}
	err := tx.Create(&gr).Error
	if err != nil {
		tx.Rollback()
	}
	for i := 0; i < len(listUserID); i++ {
		err = tx.Create(&models.GroupMember{
			GroupID: groupID,
			UserID:  listUserID[i],
		}).Error
		if err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
	return gr, err
}

func (g *group) ListGroupByUserID(userID string) ([]models.Group, error) {
	var groupMembers []models.GroupMember
	err := g.db.Where("user_id = ?", userID).Find(&groupMembers).Error
	var groupIDs []string
	for i := 0; i < len(groupMembers); i++ {
		groupIDs = append(groupIDs, groupMembers[i].GroupID)
	}
	var groupInfor []models.Group
	err = g.db.Where("group_id in ?", groupIDs).Find(&groupInfor).Error
	return groupInfor, err
}

func (g *group) ListUserByGroupID(groupID string) ([]models.GroupMember, error) {
	var groupMembers []models.GroupMember
	err := g.db.Where("group_id = ?", groupID).Find(&groupMembers).Error
	return groupMembers, err
}

func (g *group) GetGroupByGroupID(groupID string) (models.Group, error) {
	var group models.Group
	err := g.db.Where("group_id = ?", groupID).First(&group).Error
	return group, err
}

func (g *group) AddMember(groupID string, listUserID []string) error {
	var groupMembers []models.GroupMember
	for i := 0; i < len(listUserID); i++ {
		groupMembers = append(groupMembers, models.GroupMember{
			GroupID: groupID,
			UserID:  listUserID[i],
		})
	}
	return g.db.Create(&groupMembers).Error
}

type GroupRepository interface {
	CreateGroup(listUserID []string, name string, avt string) (models.Group, error)
	ListGroupByUserID(userID string) ([]models.Group, error)
	ListUserByGroupID(groupID string) ([]models.GroupMember, error)
	GetGroupByGroupID(groupID string) (models.Group, error)
	AddMember(groupID string, listUserID []string) error
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &group{
		db: db,
	}
}
