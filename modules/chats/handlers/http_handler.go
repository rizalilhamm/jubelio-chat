package handlers

import (
	"jubelio.com/chat/modules/chats/usecases"
	"jubelio.com/chat/modules/chats/repositories/commands"
	"jubelio.com/chat/modules/chats/repositories/queries"
	databases "jubelio.com/chat/packages/databases"
	"github.com/gofiber/fiber/v2"
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

    // GET /v1/points/list
    api.Get("/send-message", h.SendMessage,)
}

func (h *HTTPHandler) SendMessage(c *fiber.Ctx) error {
	// data := c.Params("id")

	return utils.Response(c, nil, "Send Message Success", fiber.StatusOK)

}