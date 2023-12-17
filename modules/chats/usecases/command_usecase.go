package usecases

import (
    "context"
    "jubelio.com/chat/migrations/models"
    "jubelio.com/chat/packages/utils"

    "github.com/go-playground/validator/v10"
)

func SendMessage(ctx context.Context, chatPayload models.Chats, messagePayload models.Messages) utils.Result {
    validate := validator.New()
	var result utils.Result

    if err := validate.Struct(chatPayload); err != nil {
        return result
    }
	
    if err := validate.Struct(messagePayload); err != nil {
        return result
    }

    return result
}
