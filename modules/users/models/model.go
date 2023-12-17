package models

type Users struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username" sql:",unique" validate:"required,max=50"`
    Email    string `json:"email" sql:",unique" validate:"required,max=100"`
    Password string `json:"password" validate:"required,max=255"`
}