package models

import "time"

type Chats struct {
    ChatID    int       `json:"chat_id"`
    ChatName  string    `json:"chat_name" validate:"required,max=100"`
    CreatedAt time.Time `json:"created_at" sql:"default:current_timestamp"`
}

