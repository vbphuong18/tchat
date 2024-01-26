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

func (g *group) ListGroupByUserID(userID string) ([]domain.Group, error) {
	groupInfor, err := g.groupRepository.ListGroupByUserID(userID)
	if err != nil {
		return nil, err
	}
	var groupInf []domain.Group
	for i := 0; i < len(groupInfor); i++ {
		groupInf = append(groupInf, domain.Group{
			GroupID: groupInfor[i].GroupID,
			Name:    groupInfor[i].Name,
			Avt:     groupInfor[i].Avt,
		})
	}
	return groupInf, nil
}

func (g *group) GetGroupByGroupID(groupID string) (domain.Group, error) {
	groupModel, err := g.groupRepository.GetGroupByGroupID(groupID)
	if err != nil {
		return domain.Group{}, err
	}
	var userIDs []string
	members, _ := g.groupRepository.ListUserByGroupID(groupID)
	for i := 0; i < len(members); i++ {
		userIDs = append(userIDs, members[i].UserID)
	}
	return domain.Group{
		GroupID:    groupModel.GroupID,
		Name:       groupModel.Name,
		Avt:        groupModel.Avt,
		ListUserID: userIDs,
	}, nil
}

func (g *group) AddMember(groupID string, listUserID []string) error {
	return g.groupRepository.AddMember(groupID, listUserID)
}

type GroupService interface {
	CreateGroup(listUserID []string, name string, avt string) (domain.Group, error)
	ListGroupByUserID(userID string) ([]domain.Group, error)
	GetGroupByGroupID(groupID string) (domain.Group, error)
	AddMember(groupID string, listUserID []string) error
}

func NewGroupService(groupRepository repositories.GroupRepository, userRepository repositories.UserRepository) GroupService {
	return &group{
		groupRepository: groupRepository,
		userRepository:  userRepository,
	}
}
