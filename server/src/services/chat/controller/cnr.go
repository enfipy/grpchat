package controller

import (
	"time"

	"github.com/enfipy/grpchat/server/src/config"
	"github.com/enfipy/grpchat/server/src/models"
	"github.com/enfipy/grpchat/server/src/services/chat"
)

type chatController struct {
	config      *config.Config
	chatUsecase chat.Usecase
}

func NewController(cnfg *config.Config, chatUsecase chat.Usecase) chat.Controller {
	return &chatController{
		config:      cnfg,
		chatUsecase: chatUsecase,
	}
}

func (cnr *chatController) UserJoin(username string) func() {
	cnr.chatUsecase.UserJoin(username)
	return func() {
		cnr.chatUsecase.UserLeft(username)
	}
}

func (cnr *chatController) GetMessages() <-chan *models.Message {
	outgoing := make(chan *models.Message)
	go func() {
		defer close(outgoing)

		newMessages := cnr.chatUsecase.SubscribeForMessages()
		for {
			newMessage, ok := <-newMessages
			if !ok {
				return
			}
			outgoing <- newMessage
		}
	}()
	go func() {
		oldMessages := cnr.chatUsecase.GetMessages(uint32(cnr.config.ChatLength))
		for oldMessage := range oldMessages {
			outgoing <- oldMessage
		}
	}()
	return outgoing
}

func (cnr *chatController) SendMessage(username, message string) {
	msg := &models.Message{
		Username:  username,
		Data:      message,
		CreatedAt: time.Now().Unix(),
	}
	cnr.chatUsecase.SaveMessage(msg)
}
