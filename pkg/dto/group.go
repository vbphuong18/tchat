package dto

type CreateGroupResponse struct {
	StatusError
}

type CreateGroupRequest struct {
	ListUser []string `json:"list_user" validate:"required,gt=2,dive,required"`
	Name     string   `json:"name"`
	Avt      string   `json:"avt"`
}

type Group struct {
	GroupID  string   `json:"group_id"`
	Name     string   `json:"name"`
	Avt      string   `json:"avt"`
	ListUser []string `json:"list_user"`
}
