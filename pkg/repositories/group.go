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

type GroupRepository interface {
	CreateGroup(listUserID []string, name string, avt string) (models.Group, error)
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &group{
		db: db,
	}
}
