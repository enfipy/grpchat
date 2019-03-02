package chat

import "github.com/enfipy/grpchat/server/src/models"

type Controller interface {
	UserJoin(username string) func()
	MessageStream(username string, incoming <-chan string) <-chan *models.Message
}
