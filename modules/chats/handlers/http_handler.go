package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"jubelio.com/chat/modules/chats/models"
	"jubelio.com/chat/modules/chats/repositories/commands"
	"jubelio.com/chat/modules/chats/repositories/queries"
	"jubelio.com/chat/modules/chats/usecases"
	databases "jubelio.com/chat/packages/databases"
	"jubelio.com/chat/packages/utils"
)

type HTTPHandler struct {
	commandUsecase usecases.CommandUsecase
	queryUsecase   usecases.QueryUsecase
}

func New() *HTTPHandler {

	postgreDb := databases.InitPostgre()

	queryPostgre := queries.NewPostgreQuery(postgreDb)
	commandPostgre := commands.NewPostgreCommand(postgreDb)
	commandUsecase := usecases.NewUsecase(commandPostgre, queryPostgre)
	queryUsecase := usecases.NewUsecase(commandPostgre, queryPostgre)

	return &HTTPHandler{
		commandUsecase: commandUsecase,
		queryUsecase:   queryUsecase,
	}
}

func (h *HTTPHandler) Mount(app *fiber.App) {
	api := app.Group("/v1/chats")

	api.Post("/messages", h.SendMessage)
	// api.Get("/messages", )
}

func (h *HTTPHandler) SendMessage(c *fiber.Ctx) error {

	var chatPayload = new(models.Chats)
	var messagesPayload = new(models.Messages)
	ctx := context.Background()

	result := h.commandUsecase.SendMessage(ctx, *chatPayload, *messagesPayload)
	
	if result.Error != nil {
		return utils.Response(c, result.Error, "Failed to send message", fiber.StatusInternalServerError)
	}

	return utils.Response(c, nil, "Send Message Success", fiber.StatusOK)
}

func (h *HTTPHandler) GetChatHistory(c *fiber.Ctx) error {
	chatPayload := new(models.RequestChat)
	ctx := context.Background()

	result := h.queryUsecase.GetChatHistory(ctx, *chatPayload)

	if result.Error != nil {
		return utils.Response(c, result.Error, "Failed to  GetChatHistory", fiber.StatusInternalServerError)
	}
	return utils.Response(c, nil, "GetChatHistory Success", fiber.StatusOK)
}
