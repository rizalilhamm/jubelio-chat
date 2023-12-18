package usecases

import (
	"context"
	"encoding/json"
	"reflect"

	"jubelio.com/chat/modules/chats/repositories/queries"

	"jubelio.com/chat/modules/chats/models"
	"jubelio.com/chat/packages/utils"
)


func (u usecase) FetchMessage(ctx context.Context, payload models.Chats) utils.Result {
	var result utils.Result

	return result
}

func (u usecase) GetChatHistory(ctx context.Context, payload models.RequestChat) utils.Result {
	var result utils.Result

	queryChatPayload := queries.QueryPayload{
		Table: "chats",
		Select: "chat_id",
		Query: "chat_name = @chat_name",
		Parameter: map[string]interface{}{"chat_name": payload.ChatName},
		Order: "sent_at ASC",
	}

	chatDetail := <-u.postgreQuery.FindOne(&queryChatPayload)

    if reflect.ValueOf(chatDetail.Data).IsNil() {
        result.Data = "Data Chat Not Found"
        return result
    } else if !reflect.ValueOf(chatDetail.Data).IsNil() {
		var chat models.Chats
		byteChat, _ := json.Marshal(chatDetail.Data)
        _ = json.Unmarshal(byteChat, &chat)
		queryMessagesPayload := queries.QueryPayload{
			Table: "messages",
			Select: "message_text",
			Query: "chat_id = @chat_id",
			Parameter: map[string]interface{}{"chat_id": chat.ChatID},
		}

		messagesResp := <-u.postgreQuery.FindMany(&queryMessagesPayload)
		result.Data = messagesResp.Data.([]models.Messages)

		return result
	}
	return result
}