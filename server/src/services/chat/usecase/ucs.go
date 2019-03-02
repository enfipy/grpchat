package usecase

import (
	"errors"
	"math/rand"

	"github.com/enfipy/grpchat/server/src/config"
	"github.com/enfipy/grpchat/server/src/models"
	"github.com/enfipy/grpchat/server/src/services/chat"

	"github.com/enfipy/locker"
	"github.com/gomodule/redigo/redis"
)

type chatUsecase struct {
	config *config.Config
	pool   *redis.Pool
	locker *locker.Locker
}

func NewUsecase(cnfg *config.Config, pool *redis.Pool, locker *locker.Locker) chat.Usecase {
	return &chatUsecase{
		config: cnfg,
		pool:   pool,
		locker: locker,
	}
}

func (ucs *chatUsecase) UserJoin(username string) {
	conn := ucs.pool.Get()
	defer conn.Close()

	key := "user:" + username
	ucs.locker.Lock(key)
	defer ucs.locker.Unlock(key)

	ucs.isExists(username)

	_, err := conn.Do("SADD", "users", username)
	if err != nil {
		panic(err)
	}
}

func (ucs *chatUsecase) UserLeft(username string) {
	conn := ucs.pool.Get()
	defer conn.Close()

	key := "user:" + username
	ucs.locker.Lock(key)
	defer ucs.locker.Unlock(key)

	_, err := conn.Do("SREM", "users", username)
	if err != nil {
		panic(err)
	}
}

func (ucs *chatUsecase) SaveMessage(message *models.Message) {
	conn := ucs.pool.Get()
	defer conn.Close()

	key := "chat"
	ucs.locker.Lock(key)
	defer ucs.locker.Unlock(key)

	msg := message.EncodeBinary()
	_, err := conn.Do("LPUSH", key, msg)
	if err != nil {
		panic(err)
	}
	_, err = conn.Do("PUBLISH", "new-messages", msg)
	if err != nil {
		panic(err)
	}

	toTrim := rand.Intn(int(ucs.config.ChatLength))
	if toTrim == 0 {
		_, err := conn.Do("LTRIM", key, 0, ucs.config.ChatLength)
		if err != nil {
			panic(err)
		}
	}
}

func (ucs *chatUsecase) GetMessages(count uint32) <-chan *models.Message {
	messages := make(chan *models.Message)
	go func() {
		defer close(messages)

		conn := ucs.pool.Get()
		defer conn.Close()

		key := "chat"
		ucs.locker.RLock(key)
		defer ucs.locker.RUnlock(key)

		res, err := redis.ByteSlices(conn.Do("LRANGE", key, 0, count-1))
		if err != nil {
			panic(err)
		}

		for _, message := range res {
			msg := new(models.Message)
			msg.DecodeBinary(string(message))
			messages <- msg
		}
	}()
	return messages
}

func (ucs *chatUsecase) SubscribeForMessages() <-chan *models.Message {
	messages := make(chan *models.Message)
	go func() {
		defer close(messages)

		conn := ucs.pool.Get()
		defer conn.Close()

		psc := redis.PubSubConn{Conn: conn}
		err := psc.Subscribe("new-messages")
		if err != nil {
			panic(err)
		}

		for {
			switch val := psc.Receive().(type) {
			case redis.Message:
				msg := new(models.Message)
				msg.DecodeBinary(string(val.Data))
				messages <- msg
			case error:
				return
			}
		}
	}()
	return messages
}

func (ucs *chatUsecase) isExists(username string) {
	conn := ucs.pool.Get()
	defer conn.Close()

	key := "chat"
	ucs.locker.RLock(key)
	defer ucs.locker.RUnlock(key)

	res, err := conn.Do("SISMEMBER", "users", username)
	if err != nil {
		panic(err)
	}

	if res.(int64) == 1 {
		panic(errors.New("Username already taken"))
	}
}
