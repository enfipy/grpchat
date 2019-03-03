package chat

import "github.com/enfipy/grpchat/server/src/models"

type Controller interface {
	UserJoin(username string) func()
	GetMessages() <-chan *models.Message
	SendMessage(username, message string)
}
