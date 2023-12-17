package models

import "time"

type Messages struct {
    MessageID   int       `json:"message_id"`
    ChatID      int       `json:"chat_id"`
    SenderID    int       `json:"sender_id"`
    ReceiverID  int       `json:"receiver_id"`
    MessageText string    `json:"message_text"`
    SentAt      time.Time `json:"sent_at" sql:"default:current_timestamp"`
}
