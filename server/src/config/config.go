package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppEnv     string
	ServerPort string

	RedisAddress string
	RedisNetwork string

	ChatLength int
}

func InitConfig() *Config {
	instance := new(Config)

	instance.AppEnv = os.Getenv("APP_ENV")
	instance.ServerPort = os.Getenv("SERVER_PORT")

	instance.RedisAddress = os.Getenv("REDIS_ADDRESS")
	instance.RedisNetwork = os.Getenv("REDIS_NETWORK")

	chatLength, err := strconv.Atoi(os.Getenv("CHAT_LENGTH"))
	if err == nil && chatLength >= 10 {
		instance.ChatLength = chatLength
	} else {
		instance.ChatLength = 100
	}

	return instance
}
