package commands

import (
	"jubelio.com/chat/packages/utils"
)

type Postgre interface {
	Create(payload *CommandPayload) <-chan utils.Result
}
