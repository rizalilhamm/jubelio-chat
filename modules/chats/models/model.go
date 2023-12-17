package models

import "time"

// Chats table struct
type Chats struct {
    ChatID    int       `json:"chat_id"`
    ChatName  string    `json:"chat_name" validate:"required,max=100"`
    CreatedAt time.Time `json:"created_at" sql:"default:current_timestamp"`
}

// Messages table struct
type Messages struct {
    MessageID   int       `json:"message_id"`
    ChatID      int       `json:"chat_id"`
    SenderID    int       `json:"sender_id" validate:"required"`
    ReceiverID  int       `json:"receiver_id" validate:"required"`
    MessageText string    `json:"message_text" validate:"required"`
    SentAt      time.Time `json:"sent_at" sql:"default:current_timestamp"`
}
