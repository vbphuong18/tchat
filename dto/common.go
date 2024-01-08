package dto

// returncode: mã lỗi
// returnmassage: mô tả lỗi
type StatusError struct {
	ReturnCode    int    `json:"return_code"`
	ReturnMessage string `json:"return_message"`
}
