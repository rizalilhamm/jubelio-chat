package commands



import (
	"jubelio.com/chat/packages/utils"

	"gorm.io/gorm"
)

// PostgreCommand model
type PostgreCommand struct {
	db *gorm.DB
}

type CommandPayload struct {
	Table     string
	Query     interface{}
	Parameter map[string]interface{}
	Document  interface{}
}

// NewPostgreQuery create new Address query
func NewPostgreCommand(db *gorm.DB) *PostgreCommand {
	return &PostgreCommand{
		db: db,
	}
}

func (c *PostgreCommand) Create(payload *CommandPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		result := c.db.Debug().Table(payload.Table).Create(payload.Document)
		if result.Error != nil {
			output <- utils.Result{Error: result}
		}
	}()

	return output
}
