package usecases

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"jubelio.com/chat/modules/chats/models"
	"jubelio.com/chat/modules/chats/repositories/commands"
	"jubelio.com/chat/modules/chats/repositories/queries"
	"jubelio.com/chat/packages/utils"

	"github.com/go-playground/validator/v10"
)

func (u usecase) SendMessage(ctx context.Context, chatPayload models.Chats, messagePayload models.Messages) utils.Result {
    validate := validator.New()
	var result utils.Result

    if err := validate.Struct(chatPayload); err != nil {
        return result
    }
	
    if err := validate.Struct(messagePayload); err != nil {
        return result
    }
    findReceiver := queries.QueryPayload{
        Table: "users",
        Select: "user_id, username",
        Query: "user_id = @user_id",
        Parameter: map[string]interface{}{"user_id": messagePayload.ReceiverID},
    }

    receiver := <-u.postgreQuery.FindOne(&findReceiver)

    if reflect.ValueOf(receiver.Data).IsNil(){
        result.Data = "Data Not Found"
        return result
    }
    /* 
    NOTE:
        - Chat Name: senderID + receiverID
    */
    chatName := strconv.Itoa(messagePayload.SenderID) + strconv.Itoa(messagePayload.ReceiverID)

    findChat := queries.QueryPayload{
        Table: "chats",
        Select: "chat_id, chat_name",
        Query: "chat_name = @chat_name",
        Parameter: map[string]interface{}{"chat_name": chatName},
    }

    /*
    NOTE:
        - Jika chat tidak ditemukan, maka create baru
        - Jika chat ditemukan, maka lanjutkan dengan yang sudah ada
    */

    chatHistory := <-u.postgreQuery.FindOne(&findChat)

    var chat models.Chats
    if !reflect.ValueOf(chatHistory.Data).IsNil() && chatHistory.Error == nil {
    	byteChat, _ := json.Marshal(chatHistory.Data)
        _ = json.Unmarshal(byteChat, &chat)
        messagePayload.ChatID = chat.ChatID

    } else if reflect.ValueOf(chatHistory.Data).IsNil() && chatHistory.Error == nil {
        /*
        NOTE:
            - Create new Chat row
        */
        chatCommandPayload := commands.CommandPayload{
            Table: "chats",
            Document: chatPayload,
        }
        newChat := <-u.postgreCommand.Create(&chatCommandPayload)
        if newChat.Error != nil {
            result.Data = "Gagal membuat Chat"
            return result
        }
    	byteChat, _ := json.Marshal(chatHistory.Data)
        _ = json.Unmarshal(byteChat, &chat)
        messagePayload.ChatID = chat.ChatID
    }
    
    messagePayload.SentAt = time.Now()
    messageCommandPayload := commands.CommandPayload{Table: "messages", Document: messagePayload,}

    newMessage := <-u.postgreCommand.Create(&messageCommandPayload)
    if newMessage.Error != nil {
        result.Data = "Gagal Send Message"
        return result
    }
    result.Data = "Berhasil Running SendMessage Method"
    return result
}
