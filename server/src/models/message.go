package models

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	chatPB "github.com/enfipy/grpchat/schema/gen/go"
)

/* Message */

type Message struct {
	Username  string
	Data      string
	CreatedAt int64
}

func (msg *Message) EncodeBinary() string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(msg)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func (msg *Message) DecodeBinary(str string) {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	decoder := gob.NewDecoder(&b)
	err = decoder.Decode(msg)
	if err != nil {
		panic(err)
	}
}

func (msg *Message) ToProtobuf() *chatPB.ServerMessage {
	return &chatPB.ServerMessage{
		Username:  msg.Username,
		Message:   msg.Data,
		CreatedAt: msg.CreatedAt,
	}
}
