package domain

type Group struct {
	GroupID    string   `json:"group_id"`
	Name       string   `json:"name"`
	Avt        string   `json:"avt"`
	ListUserID []string `json:"list_user_id"`
}
