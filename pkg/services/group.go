package services

import (
	"TChat/pkg/domain"
	"TChat/pkg/repositories"
)

type group struct {
	groupRepository repositories.GroupRepository
	userRepository  repositories.UserRepository
}

func (g *group) CreateGroup(listUserID []string, name string, avt string) (domain.Group, error) {
	if name == "" {
		for i := 0; i < len(listUserID); i++ {
			user, err := g.userRepository.GetUserByUserID(listUserID[i])
			if err != nil {
				return domain.Group{}, err
			}
			name += user.Name
			if i != len(listUserID)-1 {
				name += ", "
			}
		}
	}
	groupModel, err := g.groupRepository.CreateGroup(listUserID, name, avt)
	if err != nil {
		return domain.Group{}, err
	}
	return domain.Group{
		GroupID:    groupModel.GroupID,
		Name:       name,
		Avt:        avt,
		ListUserID: listUserID,
	}, nil
}

type GroupService interface {
	CreateGroup(listUserID []string, name string, avt string) (domain.Group, error)
}

func NewGroupService(groupRepository repositories.GroupRepository, userRepository repositories.UserRepository) GroupService {
	return &group{
		groupRepository: groupRepository,
		userRepository:  userRepository,
	}
}
