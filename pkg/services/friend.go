package services

import (
	"TChat/pkg/domain"
	"TChat/pkg/repositories"
)

type friend struct {
	friendRepository repositories.FriendRepository
}

func (f *friend) CreateFriend(friend domain.Friend) error {
	return f.friendRepository.CreateFriend(friend)
}

func (f *friend) ListFriend(userID string) ([]domain.Friend, error) {
	friendModel, err := f.friendRepository.ListFriend(userID)
	if err != nil {
		return nil, err
	}
	var friendDomain []domain.Friend
	for i := 0; i < len(friendModel); i++ {
		friendDomain = append(friendDomain, domain.Friend{
			UserID1: friendModel[i].UserID1,
			UserID2: friendModel[i].UserID2,
		})
	}
	return friendDomain, nil

}

func (f *friend) DeleteFriend(userID1 string, userID2 string) error {
	return f.friendRepository.DeleteFriend(userID1, userID2)
}

type FriendService interface {
	CreateFriend(friend domain.Friend) error
	ListFriend(userID string) ([]domain.Friend, error)
	DeleteFriend(userID1 string, userID2 string) error
}

func NewFriendService(friendRepository repositories.FriendRepository) FriendService {
	return &friend{
		friendRepository: friendRepository,
	}
}
