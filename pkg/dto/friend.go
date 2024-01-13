package dto

type CreateFriendResponse struct {
	StatusError
}

type ListFriendResponse struct {
	StatusError
	Data []Friend `json:"data"`
}

type DeleteFriendResponse struct {
	StatusError
}

type CreateFriendRequest struct {
	UserID1 string `json:"user_id_1" validate:"required"`
	UserID2 string `json:"user_id_2" validate:"required"`
}

type Friend struct {
	UserID1 string `json:"user_id_1"`
	UserID2 string `json:"user_id_2"`
}
