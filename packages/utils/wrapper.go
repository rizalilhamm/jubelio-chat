package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Result struct {
	Data     interface{}
	MetaData interface{}
	Error    interface{}
	Count    int64
}

type BaseWrapperModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Meta    interface{} `json:"meta,omitempty"`
}


func Response(c *fiber.Ctx, data interface{}, message string, code int) error {
	success := false

	if code < fiber.StatusBadRequest {
		success = true
	}

	result := BaseWrapperModel{
		Success: success,
		Data:    data,
		Message: message,
		Code:    code,
	}

	return c.Status(code).JSON(result)
}