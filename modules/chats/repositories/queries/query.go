package queries

import (
	"jubelio.com/chat/packages/utils"
)

type Postgre interface {
	Count(Payload *QueryPayload) <-chan utils.Result
	FindOne(Payload *QueryPayload) <-chan utils.Result
	FindMany(Payload *QueryPayload) <-chan utils.Result
}
