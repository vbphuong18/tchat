package domain

type Group struct {
	GroupID  string   `json:"group_id"`
	Name     string   `json:"name"`
	Avt      string   `json:"avt"`
	ListUser []string `json:"list_user"`
}
