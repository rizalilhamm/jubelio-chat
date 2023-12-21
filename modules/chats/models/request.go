package models

type RequestChat struct {
	ChatName  string    `json:"chat_name" validate:"required,max=100"`
}