package usecases

import (
	"context"

	"jubelio.com/chat/migrations/models"
	"jubelio.com/chat/modules/chats/repositories/commands"
	"jubelio.com/chat/modules/chats/repositories/queries"
	"jubelio.com/chat/packages/utils"
)

type usecase struct {
	postgreCommand commands.Postgre
	postgreQuery   queries.Postgre
}

func NewUsecase(postgreCommand commands.Postgre, postgreQuery queries.Postgre) *usecase {
	return &usecase{
		postgreCommand: postgreCommand,
		postgreQuery:   postgreQuery,
	}
}
// CommandUsecase interface
type CommandUsecase interface {
	SendMessage(ctx context.Context, chatPayload models.Chats, messagePayload models.Messages) utils.Result
}

// QueryUsecase interface
type QueryUsecase interface {
	FetchMessage(ctx context.Context, payload models.Chats) utils.Result
	// GetChatDetail(ctx context.Context, payload models.Chats)
	// GetSearchMessage(ctx context.Context, payload models.Chats) utils.Result
}
