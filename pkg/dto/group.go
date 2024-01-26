package dto

type CreateGroupResponse struct {
	StatusError
	Group Group `json:"group"`
}

type ListGroupByUserIDResponse struct {
	StatusError
	Group []Group `json:"group"`
}

type GetGroupByGroupIDResponse struct {
	StatusError
	Group Group `json:"group"`
}

type AddMemberResponse struct {
	StatusError
}

type CreateGroupRequest struct {
	ListUserID []string `json:"list_user_id" validate:"required,gt=2,dive,required"`
	Name       string   `json:"name"`
	Avt        string   `json:"avt"`
}

type AddMemberRequest struct {
	GroupID    string   `json:"group_id" validate:"required"`
	ListUserID []string `json:"list_user_id" validate:"required"`
}

type Group struct {
	GroupID    string   `json:"group_id"`
	Name       string   `json:"name"`
	Avt        string   `json:"avt"`
	ListUserID []string `json:"list_user_id,omitempty"`
}
