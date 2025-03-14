package dto

type UserResponseDto struct {
	Id     uint   `json:"id"`
	Email  string `json:"email,omitempty"`
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Userid uint   `json:"userid,omitempty"`
}
