package dto

type CreateGroupResponse struct {
	StatusError
	Group Group `json:"group"`
}

type CreateGroupRequest struct {
	ListUserID []string `json:"list_user_id" validate:"required,gt=2,dive,required"`
	Name       string   `json:"name"`
	Avt        string   `json:"avt"`
}

type Group struct {
	GroupID    string   `json:"group_id"`
	Name       string   `json:"name"`
	Avt        string   `json:"avt"`
	ListUserID []string `json:"list_user_id"`
}
