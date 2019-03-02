package chat

import "github.com/enfipy/grpchat/server/src/models"

type Usecase interface {
	UserJoin(username string)
	UserLeft(username string)

	SaveMessage(message *models.Message)
	GetMessages(count uint32) <-chan *models.Message
	SubscribeForMessages() <-chan *models.Message
}
