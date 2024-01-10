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

type FriendService interface {
	CreateFriend(friend domain.Friend) error
}

func NewFriendService(friendRepository repositories.FriendRepository) FriendService {
	return &friend{
		friendRepository: friendRepository,
	}
}
